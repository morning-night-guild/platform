package mock

import (
	"context"

	"github.com/morning-night-guild/platform/app/core/model"
	"github.com/morning-night-guild/platform/app/core/usecase/repository"
)

var _ repository.Article = (*Article)(nil)

// Article 記事リポジトリのモック.
type Article struct {
	Articles []model.Article
	Err      error
}

// Save 記事を保存するモックメソッド.
func (a *Article) Save(ctx context.Context, article model.Article) error {
	return a.Err
}

// FindAll 記事を一覧取得するモックメソッド.
func (a *Article) FindAll(
	ctx context.Context,
	index repository.Index,
	size repository.Size,
) ([]model.Article, error) {
	return a.Articles, a.Err
}
