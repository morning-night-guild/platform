package main

import (
	"github.com/morning-night-guild/platform/internal/adapter/controller"
	"github.com/morning-night-guild/platform/internal/adapter/gateway"
	"github.com/morning-night-guild/platform/internal/driver/config"
	"github.com/morning-night-guild/platform/internal/driver/database"
	"github.com/morning-night-guild/platform/internal/driver/env"
	"github.com/morning-night-guild/platform/internal/driver/handler"
	"github.com/morning-night-guild/platform/internal/driver/interceptor"
	"github.com/morning-night-guild/platform/internal/driver/newrelic"
	"github.com/morning-night-guild/platform/internal/driver/server"
	"github.com/morning-night-guild/platform/internal/usecase/interactor/article"
	"github.com/morning-night-guild/platform/pkg/log"
)

func main() {
	env.Init()

	cfg := config.NewCore()

	db := database.NewClient()

	rdb, err := db.Of(cfg.DSN)
	if err != nil {
		log.Log().Fatal(err.Error())
	}

	articleRepo := gateway.NewArticle(rdb)

	articleShareItr := article.NewShareInteractor(articleRepo)

	articleListItr := article.NewListInteractor(articleRepo)

	ctl := controller.New()

	articleCtr := controller.NewArticle(ctl, articleShareItr, articleListItr)

	healthCtr := controller.NewHealth()

	var nr *newrelic.NewRelic

	if env.Get().IsProd() {
		nr, err = newrelic.New(cfg.NewRelicAppName, cfg.NewRelicLicense)
		if err != nil {
			log.Log().Sugar().Errorf("failed to init newrelic: %s", err)
		}
	}

	ic := interceptor.New()

	h := handler.NewConnectHandler(ic, nr, articleCtr, healthCtr)

	srv := server.NewServer(cfg.Port, h)

	srv.Run()
}
