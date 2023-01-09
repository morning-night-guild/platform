//go:build e2e
// +build e2e

package article_test

import (
	"context"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-guild/platform/e2e/helper"
	articlev1 "github.com/morning-night-guild/platform/pkg/connect/proto/article/v1"
)

func TestE2EArticleShare(t *testing.T) {
	t.Parallel()

	url := helper.GetEndpoint(t)

	t.Run("記事が共有できる", func(t *testing.T) {
		t.Parallel()

		hc := &http.Client{
			Transport: helper.NewAPIKeyTransport(t, helper.GetAPIKey(t)),
		}

		client := helper.NewClient(t, hc, url)

		req := &articlev1.ShareRequest{
			Url:         "https://www.example.com",
			Title:       "title",
			Description: "description",
			Thumbnail:   "https://www.example.com/thumbnail.jpg",
		}

		got, err := client.Article.Share(context.Background(), connect.NewRequest(req))
		if err != nil {
			t.Fatalf("failed to article share: %s", err)
		}

		if !reflect.DeepEqual(got.Msg.Article.Url, req.Url) {
			t.Errorf("Url = %v, want %v", got.Msg.Article.Url, req.Url)
		}
		if !reflect.DeepEqual(got.Msg.Article.Title, req.Title) {
			t.Errorf("Title = %v, want %v", got.Msg.Article.Title, req.Title)
		}
		if !reflect.DeepEqual(got.Msg.Article.Description, req.Description) {
			t.Errorf("Description = %v, want %v", got.Msg.Article.Description, req.Description)
		}
		if !reflect.DeepEqual(got.Msg.Article.Thumbnail, req.Thumbnail) {
			t.Errorf("Thumbnail = %v, want %v", got.Msg.Article.Thumbnail, req.Thumbnail)
		}

	})

	t.Run("不正なURLが指定されて記事が共有できない", func(t *testing.T) {
		t.Parallel()

		hc := &http.Client{
			Transport: helper.NewAPIKeyTransport(t, helper.GetAPIKey(t)),
		}

		client := helper.NewClient(t, hc, url)

		req := &articlev1.ShareRequest{
			Url:         "http://www.example.com",
			Title:       "title",
			Description: "description",
			Thumbnail:   "https://www.example.com/thumbnail.jpg",
		}

		_, err := client.Article.Share(context.Background(), connect.NewRequest(req))
		if !strings.Contains(err.Error(), "invalid_argument") {
			t.Errorf("err = %v", err)
		}
		if !strings.Contains(err.Error(), "bad request") {
			t.Errorf("err = %v", err)
		}
	})

	t.Run("不正なThumbnailが指定されて記事が共有できない", func(t *testing.T) {
		t.Parallel()

		hc := &http.Client{
			Transport: helper.NewAPIKeyTransport(t, helper.GetAPIKey(t)),
		}

		client := helper.NewClient(t, hc, url)

		req := &articlev1.ShareRequest{
			Url:         "https://www.example.com",
			Title:       "title",
			Description: "description",
			Thumbnail:   "http://www.example.com/thumbnail.jpg",
		}

		_, err := client.Article.Share(context.Background(), connect.NewRequest(req))
		if !strings.Contains(err.Error(), "invalid_argument") {
			t.Errorf("err = %v", err)
		}
		if !strings.Contains(err.Error(), "bad request") {
			t.Errorf("err = %v", err)
		}
	})

	t.Run("Api-Keyがなくて記事が共有できない", func(t *testing.T) {
		t.Parallel()

		hc := &http.Client{}

		client := helper.NewClient(t, hc, url)

		req := &articlev1.ShareRequest{
			Url: "https://www.example.com",
		}

		_, err := client.Article.Share(context.Background(), connect.NewRequest(req))
		if !strings.Contains(err.Error(), "unauthenticated") {
			t.Errorf("err = %v", err)
		}
		if !strings.Contains(err.Error(), "unauthorized") {
			t.Errorf("err = %v", err)
		}
	})
}
