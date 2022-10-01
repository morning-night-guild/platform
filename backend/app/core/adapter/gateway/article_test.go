package gateway_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/morning-night-guild/platform/app/core/adapter/gateway"
	"github.com/morning-night-guild/platform/app/core/model"
	"github.com/morning-night-guild/platform/app/core/model/article"
	"github.com/morning-night-guild/platform/pkg/ent"
	"github.com/morning-night-guild/platform/pkg/ent/articletag"
	"github.com/morning-night-guild/platform/pkg/ent/enttest"
)

var _ gateway.RDBFactory = (*RDBClientMock)(nil)

type RDBClientMock struct {
	t *testing.T
}

func NewRDBClientMock(t *testing.T) *RDBClientMock {
	t.Helper()

	return &RDBClientMock{
		t: t,
	}
}

func (r *RDBClientMock) Of(dsn string) (*gateway.RDB, error) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(r.t.Log)),
	}

	dataSourceName := fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", dsn)

	db := enttest.Open(r.t, "sqlite3", dataSourceName, opts...)

	return &gateway.RDB{
		Client: db,
	}, nil
}

func TestArticleSave(t *testing.T) {
	t.Parallel()

	t.Run("記事を保存できる", func(t *testing.T) {
		t.Parallel()

		rdb, err := NewRDBClientMock(t).Of(uuid.NewString())
		if err != nil {
			t.Fatal(err)
		}

		ag := gateway.NewArticle(rdb)

		ctx := context.Background()

		a := model.CreateArticle(
			article.Title("タイトル"),
			article.URL("https://example.com"),
			article.Description("説明"),
			article.Thumbnail("https://example.com"),
			article.TagList{},
		)

		if err := ag.Save(ctx, a); err != nil {
			t.Error(err)
		}

		found, err := rdb.Article.Get(ctx, a.ID.Value())
		if err != nil {
			t.Error(err)
		}

		got, _ := model.NewArticle(
			article.ID(found.ID),
			article.Title(found.Title),
			article.URL(found.URL),
			article.Description(found.Description),
			article.Thumbnail(found.Thumbnail),
			article.TagList{},
		)

		if !reflect.DeepEqual(got, a) {
			t.Errorf("NewArticle() = %v, want %v", got, a)
		}

		// 同じURLを保存してもerrorにならないことを確認
		if err := ag.Save(ctx, model.CreateArticle(
			article.Title("タイトル"),
			article.URL("https://example.com"),
			article.Description("説明"),
			article.Thumbnail("https://example.com"),
			article.TagList{},
		)); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("タグを含む記事が保存できる", func(t *testing.T) {
		t.Parallel()

		rdb, err := NewRDBClientMock(t).Of(uuid.NewString())
		if err != nil {
			t.Fatal(err)
		}

		ag := gateway.NewArticle(rdb)

		ctx := context.Background()

		if err := ag.Save(ctx, model.CreateArticle(
			article.Title("タイトル"),
			article.URL("https://example.com"),
			article.Description("説明"),
			article.Thumbnail("https://example.com"),
			article.TagList([]article.Tag{
				article.Tag("tag1"),
				article.Tag("tag2"),
				article.Tag("tag3"),
				article.Tag("tag4"),
				article.Tag("tag5"),
			}),
		)); err != nil {
			t.Error(err)
		}
	})

	t.Run("既にある記事に既にあるタグを保存しようとしてもエラーにならない", func(t *testing.T) {
		t.Parallel()

		rdb, err := NewRDBClientMock(t).Of(uuid.NewString())
		if err != nil {
			t.Fatal(err)
		}

		ag := gateway.NewArticle(rdb)

		ctx := context.Background()

		a1 := model.CreateArticle(
			article.Title("タイトル"),
			article.URL("https://example.com1"),
			article.Description("説明"),
			article.Thumbnail("https://example.com"),
			article.TagList([]article.Tag{
				article.Tag("tag1"),
				article.Tag("tag2"),
			}),
		)

		a2 := model.CreateArticle(
			article.Title("タイトル"),
			article.URL("https://example.com1"),
			article.Description("説明"),
			article.Thumbnail("https://example.com"),
			article.TagList([]article.Tag{
				article.Tag("tag1"),
				article.Tag("tag2"),
			}),
		)

		if err := ag.Save(ctx, a1); err != nil {
			t.Fatal(err)
		}

		if err := ag.Save(ctx, a2); err != nil {
			t.Fatal(err)
		}

		found, err := rdb.ArticleTag.Query().
			Where(articletag.ArticleIDEQ(a1.ID.Value())).
			All(ctx)
		if err != nil {
			t.Fatal(err)
		}

		var got article.TagList
		for _, item := range found {
			got = append(got, article.Tag(item.Tag))
		}

		if !reflect.DeepEqual(a1.TagList, got) {
			t.Errorf("NewArticle() = %v, want %v", got, a1.TagList)
		}
	})
}
