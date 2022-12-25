package config

import (
	"os"

	"github.com/morning-night-guild/platform/pkg/log"
)

type Config struct {
	DSN    string
	APIKey string
}

var config Config //nolint:gochecknoglobals

func Init() {
	c := Config{
		DSN:    os.Getenv("DATABASE_URL"),
		APIKey: os.Getenv("API_KEY"),
	}

	log.Log().Sugar().Infof("config: %+v", c)

	config = c
}

func Get() Config {
	return config
}
