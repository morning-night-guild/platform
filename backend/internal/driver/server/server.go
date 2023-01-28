package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-guild/platform/internal/adapter/controller"
	"github.com/morning-night-guild/platform/internal/driver/config"
	"github.com/morning-night-guild/platform/internal/driver/interceptor"
	"github.com/morning-night-guild/platform/internal/driver/newrelic"
	"github.com/morning-night-guild/platform/pkg/connect/proto/article/v1/articlev1connect"
	"github.com/morning-night-guild/platform/pkg/connect/proto/health/v1/healthv1connect"
	"github.com/morning-night-guild/platform/pkg/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	shutdownTime      = 10
	readHeaderTimeout = 30 * time.Second
)

type HTTPServer struct {
	*http.Server
}

// NewHTTPServer
// 引数nrはnilでも動作可能（NewRelicへレポートが送信されない）.
func NewHTTPServer(
	nr *newrelic.NewRelic,
	article *controller.Article,
	health *controller.Health,
) (*HTTPServer, error) {
	ic := connect.WithInterceptors(interceptor.New())

	routes := []Route{
		NewRoute(articlev1connect.NewArticleServiceHandler(article, ic)),
		NewRoute(healthv1connect.NewHealthServiceHandler(health, ic)),
	}

	if nr != nil {
		for i, route := range routes {
			routes[i] = NewRoute(nr.Handle(route.path, route.handler))
		}
	}

	mux := NewRouter(routes...).Mux()

	allowOrigins, err := ConvertAllowOrigins(config.Get().CORSAllowOrigins)
	if err != nil {
		log.Log().Warn("failed to convert allow origins", log.ErrorField(err))

		return nil, err
	}

	cors, err := NewCORS(allowOrigins, ConvertDebugEnable(config.Get().CORSDebugEnable))
	if err != nil {
		log.Log().Warn("failed to create CORS config", log.ErrorField(err))

		return nil, err
	}

	s := &http.Server{
		Addr:              fmt.Sprintf(":%s", config.Get().Port),
		Handler:           cors.Handler(h2c.NewHandler(mux, &http2.Server{})),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return &HTTPServer{
		Server: s,
	}, nil
}

func (s *HTTPServer) Run() {
	log.Log().Sugar().Infof("server running on %s", s.Addr)

	go func() {
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Log().Sugar().Errorf("server closed with error: %s", err.Error())

			log.Log().Panic("server shutdown")
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	log.Log().Sugar().Infof("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime*time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Log().Sugar().Infof("failed to gracefully shutdown:", err)
	}

	log.Log().Info("HTTPServer shutdown")
}
