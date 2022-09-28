package article

import (
	"context"

	"github.com/morning-night-guild/platform/app/core/usecase/port"
	"github.com/morning-night-guild/platform/app/core/usecase/repository"
)

var _ port.ShareArticle = (*ShareInteractor)(nil)

// ShareInteractor 記事共有のインタラクター.
type ShareInteractor struct {
	articleRepository repository.Article // 記事のリポジトリ
	ogpRepository     repository.OGP     // OGPリポジトリ
}

// NewShareInteractor 記事共有のインタラクターのファクトリ関数.
func NewShareInteractor(
	articleRepository repository.Article,
	ogpRepository repository.OGP,
) *ShareInteractor {
	return &ShareInteractor{
		articleRepository: articleRepository,
		ogpRepository:     ogpRepository,
	}
}

// Execute 記事共有のインタラクターを実行する.
func (s *ShareInteractor) Execute(ctx context.Context, input port.ShareArticleInput) (port.ShareArticleOutput, error) {
	a, err := s.ogpRepository.Create(ctx, input.URL)
	if err != nil {
		return port.ShareArticleOutput{}, err
	}

	if err := s.articleRepository.Save(ctx, a); err != nil {
		return port.ShareArticleOutput{}, err
	}

	return port.ShareArticleOutput{}, nil
}
