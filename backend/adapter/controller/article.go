package controller

import (
	"context"
	"os"

	"github.com/bufbuild/connect-go"
	articlev1 "github.com/morning-night-guild/platform/driver/connect/article/v1"
	"github.com/morning-night-guild/platform/model/article"
	"github.com/morning-night-guild/platform/usecase/port"
)

// Article.
type Article struct {
	share port.ShareArticle
}

// NewArticle 記事のコントローラを新規作成する関数.
func NewArticle(
	share port.ShareArticle,
) *Article {
	return &Article{
		share: share,
	}
}

// Share 記事を共有するコントローラメソッド.
func (a Article) Share(
	ctx context.Context,
	req *connect.Request[articlev1.ShareRequest],
) (*connect.Response[articlev1.ShareResponse], error) {
	if req.Header().Get("X-API-KEY") != os.Getenv("API_KEY") {
		return nil, ErrUnauthorized
	}

	url, err := article.NewURL(req.Msg.Url)
	if err != nil {
		return nil, handleError(err)
	}

	input := port.ShareArticleInput{
		URL: url,
	}

	if _, err := a.share.Execute(ctx, input); err != nil {
		return nil, handleError(err)
	}

	return connect.NewResponse(&articlev1.ShareResponse{}), nil
}
