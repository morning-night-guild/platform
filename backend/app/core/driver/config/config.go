package config

import "os"

type Config struct {
	DSN    string
	APIKey string
}

var config Config //nolint:gochecknoglobals

func Init() {
	config = Config{
		DSN:    os.Getenv("DATABASE_URL"),
		APIKey: os.Getenv("API_KEY"),
	}
}

func Get() Config {
	return config
}
