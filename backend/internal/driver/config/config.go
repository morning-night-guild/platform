package config

import (
	"os"
	"strconv"

	"github.com/morning-night-guild/platform/pkg/log"
)

type Config struct {
	DSN              string
	APIKey           string
	Port             string
	NewRelicAppName  string
	NewRelicLicense  string
	CORSAllowOrigins string
	CORSDebugEnable  string
}

var config Config //nolint:gochecknoglobals

func Init() {
	port := os.Getenv("PORT")

	if _, err := strconv.Atoi(port); err != nil {
		port = "8080"
	}

	c := Config{
		DSN:              os.Getenv("DATABASE_URL"),
		APIKey:           os.Getenv("API_KEY"),
		Port:             port,
		NewRelicAppName:  os.Getenv("NEWRELIC_APP_NAME"),
		NewRelicLicense:  os.Getenv("NEWRELIC_LICENSE"),
		CORSAllowOrigins: os.Getenv("CORS_ALLOW_ORIGINS"),
		CORSDebugEnable:  os.Getenv("CORS_DEBUG_ENABLE"),
	}

	log.Log().Sugar().Infof("config: %+v", c)

	config = c
}

func Get() Config {
	return config
}
