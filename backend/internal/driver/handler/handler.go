package handler

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/morning-night-guild/platform/internal/adapter/controller"
	"github.com/morning-night-guild/platform/internal/driver/middleware"
	"github.com/morning-night-guild/platform/internal/driver/newrelic"
	"github.com/morning-night-guild/platform/internal/driver/router"
	"github.com/morning-night-guild/platform/pkg/connect/proto/article/v1/articlev1connect"
	"github.com/morning-night-guild/platform/pkg/connect/proto/health/v1/healthv1connect"
	"github.com/morning-night-guild/platform/pkg/openapi"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewConnectHandler(
	interceptor connect.UnaryInterceptorFunc,
	nr *newrelic.NewRelic,
	article *controller.Article,
	health *controller.Health,
) http.Handler {
	ic := connect.WithInterceptors(interceptor)

	routes := []router.Route{
		router.NewRoute(articlev1connect.NewArticleServiceHandler(article, ic)),
		router.NewRoute(healthv1connect.NewHealthServiceHandler(health, ic)),
	}

	if nr != nil {
		for i, route := range routes {
			routes[i] = router.NewRoute(nr.Handle(route.Path, route.Handler))
		}
	}

	mux := router.New(routes...).Mux()

	return h2c.NewHandler(mux, &http2.Server{})
}

const (
	baseURL = "/api"
	maxAge  = 300
)

func NewOpenAPIHandler(
	si openapi.ServerInterface,
	middleware *middleware.Middleware,
) http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           maxAge,
	}))

	return openapi.HandlerWithOptions(si, openapi.ChiServerOptions{
		BaseURL:     baseURL,
		BaseRouter:  router,
		Middlewares: []openapi.MiddlewareFunc{middleware.Handle},
	})
}
