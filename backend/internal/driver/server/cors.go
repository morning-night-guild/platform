package server

import (
	"errors"
	"strings"

	"github.com/rs/cors"
)

var (
	errEmptyArray  = errors.New("empty array")
	errEmptyString = errors.New("empty string")
)

func NewCORS(allowOrigins []string, debug bool) (*cors.Cors, error) {
	if len(allowOrigins) == 0 {
		return nil, errEmptyArray
	}

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
	}), nil
}

func ConvertAllowOrigins(allowOrigins string) ([]string, error) {
	if allowOrigins == "" {
		return nil, errEmptyString
	}

	return strings.Split(allowOrigins, ","), nil
}

func ConvertDebugEnable(debug string) bool {
	return debug == "true"
}
