package port

import (
	"github.com/morning-night-guild/platform/app/core/model/article"
	"github.com/morning-night-guild/platform/app/core/usecase"
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
