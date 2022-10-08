package middleware

import (
	"net/http"

	"github.com/morning-night-guild/platform/pkg/log"
)

func Handle(pattern string, next http.Handler) (string, http.Handler) {
	return pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(log.SetLogCtx(r.Context()))

		next.ServeHTTP(w, r)
	})
}
