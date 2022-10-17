package mock

import (
	"context"

	"github.com/morning-night-guild/platform/app/core/model"
	"github.com/morning-night-guild/platform/app/core/model/article"
	"github.com/morning-night-guild/platform/app/core/usecase/repository"
)

var _ repository.OGP = (*OGP)(nil)

// OGP OGPリポジトリのモック.
type OGP struct {
	Article model.Article
	Err     error
}

// Create 記事を作成するモックメソッド.
func (o *OGP) Create(ctx context.Context, url article.URL) (model.Article, error) {
	if o.Err != nil {
		return model.Article{}, o.Err
	}

	o.Article.URL = url

	return o.Article, nil
}
