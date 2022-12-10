package main

import (
	"os"

	"github.com/morning-night-guild/platform/app/core/adapter/controller"
	"github.com/morning-night-guild/platform/app/core/adapter/gateway"
	"github.com/morning-night-guild/platform/app/core/driver/database"
	"github.com/morning-night-guild/platform/app/core/driver/server"
	"github.com/morning-night-guild/platform/app/core/usecase/interactor/article"
)

func main() {
	db := database.NewClient()

	dsn := os.Getenv("DATABASE_URL")

	rdb, _ := db.Of(dsn)

	articleRepo := gateway.NewArticle(rdb)

	ogpRepo := gateway.NewOGP()

	articleShareItr := article.NewShareInteractor(articleRepo, ogpRepo)

	articleListItr := article.NewListInteractor(articleRepo)

	articleCtr := controller.NewArticle(articleShareItr, articleListItr)

	healthCtr := controller.NewHealth()

	server := server.NewHTTPServer(articleCtr, healthCtr)

	server.Run()
}
