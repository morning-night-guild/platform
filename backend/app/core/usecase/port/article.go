package port

import (
	"github.com/morning-night-guild/platform/app/core/model"
	"github.com/morning-night-guild/platform/app/core/model/article"
	"github.com/morning-night-guild/platform/app/core/usecase"
	"github.com/morning-night-guild/platform/app/core/usecase/repository"
)

// ShareArticleInput.
type ShareArticleInput struct {
	usecase.Input
	URL article.URL
}

// ShareArticleOutput.
type ShareArticleOutput struct {
	usecase.Output
}

// ShareArticle.
type ShareArticle interface {
	usecase.Usecase[ShareArticleInput, ShareArticleOutput]
}

// ListArticleInput.
type ListArticleInput struct {
	usecase.Input
	Index repository.Index
	Size  repository.Size
}

// ListArticleOutput.
type ListArticleOutput struct {
	usecase.Output
	Articles []model.Article
}

// ListArticle.
type ListArticle interface {
	usecase.Usecase[ListArticleInput, ListArticleOutput]
}
