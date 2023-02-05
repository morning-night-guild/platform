package main

import (
	"context"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/morning-night-guild/platform/pkg/ent"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Printf("Failed create schema: %v", err)

		return
	}

	defer client.Close()

	ctx := context.Background()

	if err := client.Debug().Schema.Create(ctx); err != nil {
		log.Printf("Failed create schema: %v", err)

		return
	}
}
