package article

import (
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-guild/platform/pkg/connect/proto/article/v1/articlev1connect"
)

type Client struct {
	Article articlev1connect.ArticleServiceClient
}

func New(t *testing.T, client connect.HTTPClient, url string) *Client {
	t.Helper()

	c := articlev1connect.NewArticleServiceClient(
		client,
		url,
	)

	return &Client{
		Article: c,
	}
}
