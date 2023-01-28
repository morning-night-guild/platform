package server

import (
	"strings"

	"github.com/rs/cors"
)

func NewCORS(allowOrigins []string, debug bool) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: allowOrigins,
		AllowedHeaders: []string{
			"Origin",
			"Content-Length",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Authorization",
			"Connect-Protocol-Version",
		},
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowCredentials: true,
		Debug:            debug,
	})
}

func ConvertAllowOrigins(allowOrigins string) []string {
	return strings.Split(allowOrigins, ",")
}

func ConvertDebugEnable(debug string) bool {
	if debug == "true" {
		return true
	}

	return false
}
