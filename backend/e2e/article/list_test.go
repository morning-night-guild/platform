//go:build e2e
// +build e2e

package article_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-guild/platform/e2e/helper"
	articlev1 "github.com/morning-night-guild/platform/pkg/connect/proto/article/v1"
)

const articleCount = uint32(5)

func TestE2EArticleList(t *testing.T) {
	t.Parallel()

	helper.BulkInsert(t, int(articleCount+5))

	url := helper.GetEndpoint(t)

	t.Run("記事が一覧できる", func(t *testing.T) {
		t.Parallel()

		hc := &http.Client{
			Transport: helper.NewAPIKeyTransport(t, helper.GetAPIKey(t)),
		}

		client := helper.NewClient(t, hc, url)

		req := &articlev1.ListRequest{
			MaxPageSize: articleCount,
		}

		res, err := client.Article.List(context.Background(), connect.NewRequest(req))
		if err != nil {
			t.Fatalf("failed to article share: %s", err)
		}

		if !reflect.DeepEqual(len(res.Msg.Articles), int(articleCount)) {
			t.Errorf("Articles length = %v, want %v", len(res.Msg.Articles), articleCount)
		}
		if res.Msg.NextPageToken == "" {
			t.Error("Next Page Token is empty")
		}
	})
}
