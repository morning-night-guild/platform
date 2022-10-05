package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/morning-night-guild/platform/app/core/adapter/controller"
	"github.com/morning-night-guild/platform/app/core/driver/middleware"
	"github.com/morning-night-guild/platform/pkg/connect/article/v1/articlev1connect"
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
	mux := http.NewServeMux()

	mux.Handle(middleware.Handle(articlev1connect.NewArticleServiceHandler(article)))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	s := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return &HTTPServer{
		Server: s,
	}
}

func (s *HTTPServer) Run() {
	log.Printf("Server running on %s", s.Addr)

	go func() {
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("Server closed with error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime*time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Println("Failed to gracefully shutdown:", err)
	}

	log.Println("HTTPServer shutdown")
}
