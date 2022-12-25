package server

import (
	"net/http"

	"github.com/morning-night-guild/platform/app/core/driver/newrelic"
)

type Router struct {
	routes []Route
}

type Route struct {
	path    string
	handler http.Handler
}

// NewRoute.
func NewRoute(path string, handler http.Handler) Route {
	return Route{
		path:    path,
		handler: handler,
	}
}

// NewRouter.
func NewRouter(routes ...Route) *Router {
	return &Router{
		routes: routes,
	}
}

// Mux.
// 引数nrがnilであっても動作可能
// nilでなかった場合はnewrelicへレポートを送信.
func (r Router) Mux(nr *newrelic.NewRelic) *http.ServeMux {
	mux := http.NewServeMux()

	for _, route := range r.routes {
		path := route.path
		handler := route.handler

		if nr != nil {
			path, handler = nr.Handle(route.path, route.handler)
		}

		mux.Handle(path, handler)
	}

	return mux
}
