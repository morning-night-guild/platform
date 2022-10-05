package middleware

import (
	"log"
	"net/http"
)

func Handle(pattern string, next http.Handler) (string, http.Handler) {
	return pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s", pattern)
		next.ServeHTTP(w, r)
		log.Printf("%s", pattern)
	})
}
