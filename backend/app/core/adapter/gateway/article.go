package gateway

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/morning-night-guild/platform/app/core/model"
	"github.com/morning-night-guild/platform/app/core/usecase/repository"
	"github.com/morning-night-guild/platform/pkg/ent"
	"github.com/morning-night-guild/platform/pkg/ent/article"
	"github.com/morning-night-guild/platform/pkg/ent/articletag"
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

	now := time.Now().UTC()

	err := a.rdb.Article.Create().
		SetID(id).
		SetTitle(item.Title.String()).
		SetURL(item.URL.String()).
		SetDescription(item.Description.String()).
		SetThumbnail(item.Thumbnail.String()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		OnConflict().
		DoNothing().
		Exec(ctx)

	if err != nil && a.rdb.IsDuplicatedError(ctx, err) {
		if ea, err := a.rdb.Article.Query().Where(article.URLEQ(item.URL.String())).First(ctx); err == nil {
			id = ea.ID
		} else {
			return errors.Wrap(err, "failed to save")
		}
	} else if err != nil {
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

	if a.rdb.IsDuplicatedError(ctx, err) {
		return nil
	}

	return errors.Wrap(err, "failed to save")
}

// FindAll 記事を取得するメソッド.
func (a *Article) FindAll(ctx context.Context, index repository.Index, size repository.Size) ([]model.Article, error) {
	// ent articles
	eas, err := a.rdb.Article.Query().
		Order(ent.Desc(article.FieldCreatedAt)).
		Offset(index.Int()).
		Limit(size.Int()).
		All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to article query")
	}

	// article ids
	aids := make([]uuid.UUID, len(eas))

	// ent article map
	eam := make(map[uuid.UUID]ent.Article, len(eas))

	for i, ea := range eas {
		aids[i] = ea.ID
		eam[ea.ID] = *ea
	}

	// ent article tags
	eats, err := a.rdb.ArticleTag.Query().
		Where(articletag.ArticleIDIn(aids...)).
		All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to article tag query")
	}

	// each article id tag slice
	// ent article tag map
	eatm := make(map[uuid.UUID][]string)

	for _, eat := range eats {
		eatm[eat.ArticleID] = append(eatm[eat.ArticleID], eat.Tag)
	}

	articles := make([]model.Article, len(eas))

	for i, ea := range eas {
		tags, ok := eatm[ea.ID]
		if !ok {
			tags = []string{}
		}

		articles[i] = model.ReconstructArticle(
			ea.ID,
			ea.URL,
			ea.Title,
			ea.Description,
			ea.Thumbnail,
			tags,
		)
	}

	return articles, nil
}
