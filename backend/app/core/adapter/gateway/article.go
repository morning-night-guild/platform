package gateway

import (
	"context"
	"database/sql"

	"github.com/morning-night-guild/platform/app/core/model"
	"github.com/morning-night-guild/platform/app/core/usecase/repository"
	"github.com/morning-night-guild/platform/pkg/ent"
	"github.com/morning-night-guild/platform/pkg/ent/article"
	"github.com/morning-night-guild/platform/pkg/log"
	"github.com/pkg/errors"
)

var _ repository.Article = (*Article)(nil)

// Article.
type Article struct {
	rdb *RDB
}

// NewArticle ArticleGatewayを生成するファクトリー関数.
func NewArticle(rdb *RDB) *Article {
	return &Article{
		rdb: rdb,
	}
}

// Save 記事を保存するメソッド.
func (a *Article) Save(ctx context.Context, item model.Article) error {
	id := item.ID.Value()

	err := a.rdb.Article.Create().
		SetID(id).
		SetTitle(item.Title.String()).
		SetURL(item.URL.String()).
		SetDescription(item.Description.String()).
		SetThumbnail(item.Thumbnail.String()).
		OnConflict().
		DoNothing().
		Exec(ctx)

	if err != nil && isDuplicatedError(ctx, err) {
		if ea, err := a.findByURL(ctx, item.URL.String()); err == nil {
			id = ea.ID
		} else {
			return err
		}
	} else {
		return errors.Wrap(err, "failed to save")
	}

	if item.TagList.Len() == 0 {
		return nil
	}

	bulk := make([]*ent.ArticleTagCreate, item.TagList.Len())
	for i, tag := range item.TagList {
		bulk[i] = a.rdb.ArticleTag.Create().
			SetTag(tag.String()).
			SetArticleID(id)
	}

	err = a.rdb.ArticleTag.CreateBulk(bulk...).
		OnConflict().
		DoNothing().
		Exec(ctx)

	if err == nil {
		return nil
	}

	if isDuplicatedError(ctx, err) {
		return nil
	}

	return errors.Wrap(err, "failed to save")
}

// isDuplicatedError 重複エラーであるかを判定する関数.
func isDuplicatedError(ctx context.Context, err error) bool {
	// https://github.com/ent/ent/issues/2176 により、
	// on conflict do nothingとしてもerror no rowsが返るため、個別にハンドリングする
	if errors.Is(err, sql.ErrNoRows) {
		log := log.GetLogCtx(ctx)

		log.Debug(err.Error())

		return true
	}

	return false
}

// findByURL URLから記事を検索するメソッド.
func (a *Article) findByURL(ctx context.Context, url string) (*ent.Article, error) {
	return a.rdb.Article.Query().Where(article.URLEQ(url)).First(ctx)
}
