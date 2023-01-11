package main

import (
	"github.com/morning-night-guild/platform/app/core/adapter/controller"
	"github.com/morning-night-guild/platform/app/core/adapter/gateway"
	"github.com/morning-night-guild/platform/app/core/driver/config"
	"github.com/morning-night-guild/platform/app/core/driver/database"
	"github.com/morning-night-guild/platform/app/core/driver/env"
	"github.com/morning-night-guild/platform/app/core/driver/newrelic"
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

	articleShareItr := article.NewShareInteractor(articleRepo)

	articleListItr := article.NewListInteractor(articleRepo)

	articleCtr := controller.NewArticle(config.Get().APIKey, articleShareItr, articleListItr)

	healthCtr := controller.NewHealth()

	var nr *newrelic.NewRelic

	if env.Get().IsProd() {
		nr, err = newrelic.New(config.Get().NewRelicAppName, config.Get().NewRelicLicense)
		if err != nil {
			log.Log().Sugar().Errorf("failed to init newrelic: %s", err)
		}
	}

	server := server.NewHTTPServer(nr, articleCtr, healthCtr)

	server.Run()
}
