package main

import (
	"github.com/morning-night-guild/platform/internal/adapter/api"
	"github.com/morning-night-guild/platform/internal/driver/config"
	"github.com/morning-night-guild/platform/internal/driver/connect"
	"github.com/morning-night-guild/platform/internal/driver/env"
	"github.com/morning-night-guild/platform/internal/driver/handler"
	"github.com/morning-night-guild/platform/internal/driver/middleware"
	"github.com/morning-night-guild/platform/internal/driver/server"
)

func main() {
	env.Init()

	cfg := config.NewAPI()

	c, err := connect.New().Of(cfg.AppCoreURL)
	if err != nil {
		panic(err)
	}

	h := handler.NewOpenAPIHandler(api.New(c), middleware.New())

	srv := server.NewServer(cfg.Port, h)

	srv.Run()
}
