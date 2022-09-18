package repository

import (
	"context"

	"github.com/morning-night-guild/platform/model"
)

type Article interface {
	Save(context.Context, model.Article) error
}
