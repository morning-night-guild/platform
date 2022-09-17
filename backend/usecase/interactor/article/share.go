package article

import (
	"context"

	"github.com/morning-night-guild/platform/model"
	"github.com/morning-night-guild/platform/model/article"
	"github.com/morning-night-guild/platform/usecase/port"
	"github.com/morning-night-guild/platform/usecase/repository"
)

var _ port.ShareArticle = (*ShareInteractor)(nil)

// ShareInteractor 記事共有のインタラクター.
type ShareInteractor struct {
	articleRepository repository.Article // 記事のリポジトリ
}

// NewShareInteractor 記事共有のインタラクターのファクトリ関数.
func NewShareInteractor(articleRepository repository.Article) *ShareInteractor {
	return &ShareInteractor{
		articleRepository: articleRepository,
	}
}

// Execute 記事共有のインタラクターを実行する.
func (s *ShareInteractor) Execute(ctx context.Context, input port.ShareArticleInput) (port.ShareArticleOutput, error) {
	article := model.CreateArticle(
		input.Title,
		input.URL,
		input.Thumbnail,
		article.TagList{},
	)

	if err := s.articleRepository.Save(ctx, article); err != nil {
		return port.ShareArticleOutput{}, err
	}

	return port.ShareArticleOutput{}, nil
}
