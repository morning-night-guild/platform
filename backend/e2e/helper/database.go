package helper

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/morning-night-guild/platform/app/core/driver/database"
	"github.com/morning-night-guild/platform/pkg/ent"
)

func BulkInsert(t *testing.T, count int) {
	t.Helper()

	dsn := os.Getenv("DATABASE_URL")

	client, err := database.NewClient().Of(dsn)
	if err != nil {
		t.Fatalf("Failed to create database client: %s", err)
	}

	defer client.Close()

	ids := make([]string, count)

	for i := 0; i < count; i++ {
		ids[i] = fmt.Sprintf("00000000-0000-0000-0000-%012d", i)
	}

	bulk := make([]*ent.ArticleCreate, count)

	for i, id := range ids {
		bulk[i] = client.Article.Create().
			SetID(uuid.MustParse(id)).
			SetTitle("title-" + id).
			SetURL("https://example.com/" + id).
			SetDescription("description").
			SetThumbnail("https://example.com/" + id).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now())
	}

	if err := client.Article.CreateBulk(bulk...).
		OnConflict().
		UpdateNewValues().
		DoNothing().
		Exec(context.Background()); err != nil {
		t.Fatal(err)
	}
}
