package config

import (
	"os"
	"strconv"

	"github.com/morning-night-guild/platform/pkg/log"
)

type CoreConfig struct {
	Port             string
	DSN              string
	APIKey           string
	NewRelicAppName  string
	NewRelicLicense  string
	CORSAllowOrigins string
	CORSDebugEnable  string
}

func NewCore() CoreConfig {
	port := os.Getenv("PORT")

	if _, err := strconv.Atoi(port); err != nil {
		port = "8080"
	}

	c := CoreConfig{
		Port:             port,
		DSN:              os.Getenv("DATABASE_URL"),
		APIKey:           os.Getenv("API_KEY"),
		NewRelicAppName:  os.Getenv("NEWRELIC_APP_NAME"),
		NewRelicLicense:  os.Getenv("NEWRELIC_LICENSE"),
		CORSAllowOrigins: os.Getenv("CORS_ALLOW_ORIGINS"),
		CORSDebugEnable:  os.Getenv("CORS_DEBUG_ENABLE"),
	}

	log.Log().Sugar().Infof("config: %+v", c)

	return c
}
