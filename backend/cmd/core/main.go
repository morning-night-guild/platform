package main

import (
	"github.com/morning-night-guild/platform/internal/adapter/controller"
	"github.com/morning-night-guild/platform/internal/adapter/gateway"
	"github.com/morning-night-guild/platform/internal/driver/config"
	"github.com/morning-night-guild/platform/internal/driver/database"
	"github.com/morning-night-guild/platform/internal/driver/env"
	"github.com/morning-night-guild/platform/internal/driver/newrelic"
	"github.com/morning-night-guild/platform/internal/driver/server"
	"github.com/morning-night-guild/platform/internal/usecase/interactor/article"
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

	server, err := server.NewHTTPServer(nr, articleCtr, healthCtr)
	if err != nil {
		log.Log().Fatal(err.Error())
	}

	server.Run()
}
