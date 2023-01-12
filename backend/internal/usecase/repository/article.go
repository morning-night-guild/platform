package repository

import (
	"context"

	"github.com/morning-night-guild/platform/internal/model"
)

type Article interface {
	Save(context.Context, model.Article) error
	FindAll(context.Context, Index, Size) ([]model.Article, error)
}
