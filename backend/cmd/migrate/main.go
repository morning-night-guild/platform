package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/morning-night-guild/platform/pkg/ent"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Printf("Failed create schema: %v", err)

		return
	}

	drv := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv))

	defer client.Close()

	ctx := context.Background()

	if err := client.Debug().Schema.Create(ctx); err != nil {
		log.Printf("Failed create schema: %v", err)

		return
	}
}
