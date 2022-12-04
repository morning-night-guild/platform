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
	"github.com/morning-night-guild/platform/app/core/adapter/controller"
	"github.com/morning-night-guild/platform/app/core/driver/interceptor"
	"github.com/morning-night-guild/platform/pkg/connect/article/v1/articlev1connect"
	"github.com/morning-night-guild/platform/pkg/log"
	"github.com/rs/cors"
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

func NewHTTPServer(
	article *controller.Article,
) *HTTPServer {
	ic := connect.WithInterceptors(interceptor.New())

	mux := NewRouter(
		NewRoute(articlev1connect.NewArticleServiceHandler(article, ic)),
	).Mux()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	s := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           cors.Default().Handler(h2c.NewHandler(mux, &http2.Server{})),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return &HTTPServer{
		Server: s,
	}
}

func (s *HTTPServer) Run() {
	log.Log().Sugar().Infof("Server running on %s", s.Addr)

	go func() {
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Log().Sugar().Errorf("Server closed with error: %s", err.Error())

			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	log.Log().Sugar().Infof("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime*time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Log().Sugar().Infof("Failed to gracefully shutdown:", err)
	}

	log.Log().Info("HTTPServer shutdown")
}
