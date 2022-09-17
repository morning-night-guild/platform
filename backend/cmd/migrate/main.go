package main

import (
	"context"
	"log"
	"os"

	"github.com/morning-night-guild/platform/driver/database"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")

	client, err := database.NewClient().Of(dsn)
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
