package article_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/morning-night-guild/platform/e2e/helper"
	"github.com/morning-night-guild/platform/pkg/openapi"
)

func TestAppAPIE2EArticleShare(t *testing.T) {
	t.Parallel()

	url := helper.GetAppAPIEndpoint(t)

	t.Run("記事が共有できる", func(t *testing.T) {
		t.Parallel()

		title := uuid.NewString()

		db := helper.NewDatabase(t, helper.GetDSN(t))

		defer db.Close()

		defer db.DeleteArticleByTitle(title)

		client := helper.NewOpenAPIClientWithAPIKey(t, url, helper.GetAPIKey(t))

		res, err := client.Client.V1ShareArticle(context.Background(), openapi.V1ShareArticleRequest{
			Url:         "https://example.com",
			Title:       helper.ToStringPointer(title),
			Description: helper.ToStringPointer("description"),
			Thumbnail:   helper.ToStringPointer("https://example.com/thumbnail.jpg"),
		})
		if err != nil {
			t.Fatalf("failed to share article: %s", err)
		}

		defer res.Body.Close()

		if !reflect.DeepEqual(res.StatusCode, http.StatusOK) {
			t.Errorf("StatusCode = %v, want %v", res.StatusCode, http.StatusOK)
		}
	})

	t.Run("urlが空文字で記事が共有できない", func(t *testing.T) {
		t.Parallel()

		client := helper.NewOpenAPIClientWithAPIKey(t, url, helper.GetAPIKey(t))

		res, err := client.Client.V1ShareArticle(context.Background(), openapi.V1ShareArticleRequest{
			Url:         "",
			Title:       helper.ToStringPointer("title"),
			Description: helper.ToStringPointer("description"),
			Thumbnail:   helper.ToStringPointer("https://example.com/thumbnail.jpg"),
		})
		if err != nil {
			t.Fatalf("failed to share article: %s", err)
		}

		defer res.Body.Close()

		if !reflect.DeepEqual(res.StatusCode, http.StatusBadRequest) {
			t.Errorf("StatusCode = %v, want %v", res.StatusCode, http.StatusBadRequest)
		}
	})

	t.Run("Api-Keyがなくて記事が共有できない", func(t *testing.T) {
		t.Parallel()

		client := helper.NewOpenAPIClient(t, url)

		res, err := client.Client.V1ShareArticle(context.Background(), openapi.V1ShareArticleRequest{
			Url:         "",
			Title:       helper.ToStringPointer("title"),
			Description: helper.ToStringPointer("description"),
			Thumbnail:   helper.ToStringPointer("https://example.com/thumbnail.jpg"),
		})
		if err != nil {
			t.Fatalf("failed to share article: %s", err)
		}

		defer res.Body.Close()

		if !reflect.DeepEqual(res.StatusCode, http.StatusUnauthorized) {
			t.Errorf("StatusCode = %v, want %v", res.StatusCode, http.StatusUnauthorized)
		}
	})
}
