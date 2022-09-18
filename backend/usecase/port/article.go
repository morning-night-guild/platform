package port

import (
	"github.com/morning-night-guild/platform/model/article"
	"github.com/morning-night-guild/platform/usecase"
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
