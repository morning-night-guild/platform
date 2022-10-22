//go:build wireinject
// +build wireinject

package main

import (
	"os"

	"github.com/google/wire"
	"github.com/morning-night-guild/platform/app/core/adapter/controller"
	"github.com/morning-night-guild/platform/app/core/adapter/gateway"
	"github.com/morning-night-guild/platform/app/core/driver/database"
	"github.com/morning-night-guild/platform/app/core/driver/server"
	"github.com/morning-night-guild/platform/app/core/usecase/interactor/article"
	"github.com/morning-night-guild/platform/app/core/usecase/port"
	"github.com/morning-night-guild/platform/app/core/usecase/repository"
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
	wire.Bind(new(port.ListArticle), new(*article.ListInteractor)),
	article.NewListInteractor,
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
