package repository

import (
	"context"

	"github.com/morning-night-guild/platform/model"
	"github.com/morning-night-guild/platform/model/article"
)

type OGP interface {
	Create(context.Context, article.URL) (model.Article, error)
}
