package repository

import (
	"context"

	"github.com/morning-night-guild/platform/app/core/model"
	"github.com/morning-night-guild/platform/app/core/model/article"
)

type OGP interface {
	Create(context.Context, article.URL) (model.Article, error)
}
