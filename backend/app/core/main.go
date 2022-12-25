package main

import (
	"github.com/morning-night-guild/platform/app/core/adapter/controller"
	"github.com/morning-night-guild/platform/app/core/adapter/gateway"
	"github.com/morning-night-guild/platform/app/core/driver/config"
	"github.com/morning-night-guild/platform/app/core/driver/database"
	"github.com/morning-night-guild/platform/app/core/driver/env"
	"github.com/morning-night-guild/platform/app/core/driver/server"
	"github.com/morning-night-guild/platform/app/core/usecase/interactor/article"
	"github.com/morning-night-guild/platform/pkg/log"
)

func main() {
	env.Init()

	config.Init()

	db := database.NewClient()

	rdb, err := db.Of(config.Get().DSN)
	if err != nil {
		log.Log().Fatal(err.Error())
	}

	articleRepo := gateway.NewArticle(rdb)

	ogpRepo := gateway.NewOGP()

	articleShareItr := article.NewShareInteractor(articleRepo, ogpRepo)

	articleListItr := article.NewListInteractor(articleRepo)

	articleCtr := controller.NewArticle(config.Get().APIKey, articleShareItr, articleListItr)

	healthCtr := controller.NewHealth()

	server := server.NewHTTPServer(articleCtr, healthCtr)

	server.Run()
}
