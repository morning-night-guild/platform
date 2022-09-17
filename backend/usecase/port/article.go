package port

import (
	"github.com/morning-night-guild/platform/model/article"
	"github.com/morning-night-guild/platform/usecase"
)

// ShareArticleInput.
type ShareArticleInput struct {
	usecase.Input
	Title     article.Title
	URL       article.URL
	Thumbnail article.Thumbnail
}

// ShareArticleOutput.
type ShareArticleOutput struct {
	usecase.Output
}

// ShareArticle.
type ShareArticle interface {
	usecase.Usecase[ShareArticleInput, ShareArticleOutput]
}
