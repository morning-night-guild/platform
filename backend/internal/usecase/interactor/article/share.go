package article

import (
	"context"

	"github.com/morning-night-guild/platform/internal/model"
	"github.com/morning-night-guild/platform/internal/model/article"
	"github.com/morning-night-guild/platform/internal/usecase/port"
	"github.com/morning-night-guild/platform/internal/usecase/repository"
)

var _ port.ShareArticle = (*ShareInteractor)(nil)

// ShareInteractor 記事共有のインタラクター.
type ShareInteractor struct {
	articleRepository repository.Article // 記事のリポジトリ
}

// NewShareInteractor 記事共有のインタラクターのファクトリ関数.
func NewShareInteractor(
	articleRepository repository.Article,
) *ShareInteractor {
	return &ShareInteractor{
		articleRepository: articleRepository,
	}
}

// Execute 記事共有のインタラクターを実行する.
func (s *ShareInteractor) Execute(ctx context.Context, input port.ShareArticleInput) (port.ShareArticleOutput, error) {
	art := model.CreateArticle(input.URL, input.Title, input.Description, input.Thumbnail, []article.Tag{})

	if err := s.articleRepository.Save(ctx, art); err != nil {
		return port.ShareArticleOutput{}, err
	}

	return port.ShareArticleOutput{
		Article: art,
	}, nil
}
