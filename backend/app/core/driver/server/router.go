package server

import (
	"net/http"
)

type Router struct {
	routes []Route
}

type Route struct {
	path    string
	handler http.Handler
}

func NewRoute(path string, handler http.Handler) Route {
	return Route{
		path:    path,
		handler: handler,
	}
}

func NewRouter(routes ...Route) *Router {
	return &Router{
		routes: routes,
	}
}

func (r Router) Mux() *http.ServeMux {
	mux := http.NewServeMux()

	for _, route := range r.routes {
		mux.Handle(route.path, route.handler)
	}

	return mux
}
