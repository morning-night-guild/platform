//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/morning-night-guild/platform/adapter/controller"
	"github.com/morning-night-guild/platform/adapter/gateway"
	"github.com/morning-night-guild/platform/driver/database"
	"github.com/morning-night-guild/platform/driver/server"
	"github.com/morning-night-guild/platform/usecase/interactor/article"
	"github.com/morning-night-guild/platform/usecase/port"
	"github.com/morning-night-guild/platform/usecase/repository"
	"os"
)

func getRDB() *gateway.RDB {
	client := database.NewClient()

	dsn := os.Getenv("DATABASE_URL")

	rdb, _ := client.Of(dsn)

	return rdb
}

var driverSet = wire.NewSet(
	getRDB,
	server.NewHTTPServer,
)

var gatewaySet = wire.NewSet(
	wire.Bind(new(repository.Article), new(*gateway.Article)),
	gateway.NewArticle,
	wire.Bind(new(repository.OGP), new(*gateway.OGP)),
	gateway.NewOGP,
)

var interactorSet = wire.NewSet(
	wire.Bind(new(port.ShareArticle), new(*article.ShareInteractor)),
	article.NewShareInteractor,
)

var controllerSet = wire.NewSet(
	controller.NewArticle,
)

func InitializeHTTPServer() *server.HTTPServer {
	wire.Build(
		gatewaySet,
		interactorSet,
		controllerSet,
		driverSet,
	)

	return nil
}
