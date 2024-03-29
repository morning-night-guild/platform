package gateway_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/morning-night-guild/platform/internal/adapter/gateway"
	"github.com/morning-night-guild/platform/internal/domain/model"
	"github.com/morning-night-guild/platform/internal/domain/model/article"
	"github.com/morning-night-guild/platform/internal/domain/repository"
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

func (rm *RDBClientMock) Of(dsn string) (*gateway.RDB, error) {
	rm.t.Helper()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(rm.t.Log)),
	}

	dataSourceName := fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", dsn)

	db := enttest.Open(rm.t, "sqlite3", dataSourceName, opts...)

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
			article.URL("https://example.com"),
			article.Title("title"),
			article.Description("description"),
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
			article.URL(found.URL),
			article.Title(found.Title),
			article.Description(found.Description),
			article.Thumbnail(found.Thumbnail),
			article.TagList{},
		)

		if !reflect.DeepEqual(got, a) {
			t.Errorf("NewArticle() = %v, want %v", got, a)
		}

		// 同じURLを保存してもerrorにならないことを確認
		if err := ag.Save(ctx, model.CreateArticle(
			article.URL("https://example.com"),
			article.Title("title"),
			article.Description("description"),
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
			article.URL("https://example.com"),
			article.Title("title"),
			article.Description("description"),
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
			article.URL("https://example.com"),
			article.Title("title"),
			article.Description("description"),
			article.Thumbnail("https://example.com"),
			article.TagList([]article.Tag{
				article.Tag("tag1"),
				article.Tag("tag2"),
			}),
		)

		a2 := model.CreateArticle(
			article.URL("https://example.com"),
			article.Title("title"),
			article.Description("description"),
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

func TestArticleList(t *testing.T) {
	t.Parallel()

	t.Run("記事を一覧できる（単数）", func(t *testing.T) {
		t.Parallel()

		rdb, err := NewRDBClientMock(t).Of(uuid.NewString())
		if err != nil {
			t.Fatal(err)
		}

		ag := gateway.NewArticle(rdb)

		ctx := context.Background()

		item1 := model.CreateArticle(
			article.URL("https://example.com/1"),
			article.Title("title1"),
			article.Description("description"),
			article.Thumbnail("https://example.com/1"),
			article.TagList([]article.Tag{
				article.Tag("tag1"),
				article.Tag("tag2"),
			}),
		)

		if err := ag.Save(ctx, item1); err != nil {
			t.Fatal(err)
		}

		item2 := model.CreateArticle(
			article.URL("https://example.com/2"),
			article.Title("title2"),
			article.Description("description"),
			article.Thumbnail("https://example.com/2"),
			article.TagList([]article.Tag{
				article.Tag("tag1"),
				article.Tag("tag2"),
			}),
		)

		if err := ag.Save(ctx, item2); err != nil {
			t.Fatal(err)
		}

		got, err := ag.FindAll(ctx, repository.Index(0), repository.Size(1))
		if err != nil {
			t.Fatal(err)
		}

		articles := []model.Article{item2}

		if !reflect.DeepEqual(got, articles) {
			t.Errorf("FindAll() = %v, want %v", got, articles)
		}
	})

	t.Run("オフセットを指定して記事を一覧できる（単数）", func(t *testing.T) {
		t.Parallel()

		rdb, err := NewRDBClientMock(t).Of(uuid.NewString())
		if err != nil {
			t.Fatal(err)
		}

		ag := gateway.NewArticle(rdb)

		ctx := context.Background()

		item1 := model.CreateArticle(
			article.URL("https://example.com/1"),
			article.Title("title1"),
			article.Description("description"),
			article.Thumbnail("https://example.com/1"),
			article.TagList([]article.Tag{
				article.Tag("tag1"),
				article.Tag("tag2"),
			}),
		)

		if err := ag.Save(ctx, item1); err != nil {
			t.Fatal(err)
		}

		item2 := model.CreateArticle(
			article.URL("https://example.com/2"),
			article.Title("title2"),
			article.Description("description"),
			article.Thumbnail("https://example.com/2"),
			article.TagList([]article.Tag{
				article.Tag("tag1"),
				article.Tag("tag2"),
			}),
		)

		if err := ag.Save(ctx, item2); err != nil {
			t.Fatal(err)
		}

		got, err := ag.FindAll(ctx, repository.Index(1), repository.Size(1))
		if err != nil {
			t.Fatal(err)
		}

		articles := []model.Article{item1}

		if !reflect.DeepEqual(got, articles) {
			t.Errorf("FindAll() = %v, want %v", got, articles)
		}
	})

	t.Run("記事を一覧できる（複数）", func(t *testing.T) {
		t.Parallel()

		rdb, err := NewRDBClientMock(t).Of(uuid.NewString())
		if err != nil {
			t.Fatal(err)
		}

		ag := gateway.NewArticle(rdb)

		ctx := context.Background()

		item1 := model.CreateArticle(
			article.URL("https://example.com/1"),
			article.Title("title1"),
			article.Description("description"),
			article.Thumbnail("https://example.com/1"),
			article.TagList([]article.Tag{
				article.Tag("tag1"),
				article.Tag("tag2"),
			}),
		)

		if err := ag.Save(ctx, item1); err != nil {
			t.Fatal(err)
		}

		item2 := model.CreateArticle(
			article.URL("https://example.com/2"),
			article.Title("title2"),
			article.Description("description"),
			article.Thumbnail("https://example.com/2"),
			article.TagList([]article.Tag{
				article.Tag("tag1"),
				article.Tag("tag2"),
			}),
		)

		if err := ag.Save(ctx, item2); err != nil {
			t.Fatal(err)
		}

		got, err := ag.FindAll(ctx, repository.Index(0), repository.Size(2))
		if err != nil {
			t.Fatal(err)
		}

		articles := []model.Article{item2, item1}

		if !reflect.DeepEqual(got, articles) {
			t.Errorf("FindAll() = %v, want %v", got, articles)
		}
	})

	t.Run("保存されている記事数を超えるサイズを指定して記事を一覧できる", func(t *testing.T) {
		t.Parallel()

		rdb, err := NewRDBClientMock(t).Of(uuid.NewString())
		if err != nil {
			t.Fatal(err)
		}

		ag := gateway.NewArticle(rdb)

		ctx := context.Background()

		item := model.CreateArticle(
			article.URL("https://example.com"),
			article.Title("title"),
			article.Description("description"),
			article.Thumbnail("https://example.com"),
			article.TagList([]article.Tag{
				article.Tag("tag1"),
				article.Tag("tag2"),
			}),
		)

		if err := ag.Save(ctx, item); err != nil {
			t.Fatal(err)
		}

		got, err := ag.FindAll(ctx, repository.Index(0), repository.Size(2))
		if err != nil {
			t.Fatal(err)
		}

		articles := []model.Article{item}

		if !reflect.DeepEqual(got, articles) {
			t.Errorf("FindAll() = %v, want %v", got, articles)
		}
	})

	t.Run("保存されている記事数を超えてインデックスを指定して記事を一覧できる", func(t *testing.T) {
		t.Parallel()

		rdb, err := NewRDBClientMock(t).Of(uuid.NewString())
		if err != nil {
			t.Fatal(err)
		}

		ag := gateway.NewArticle(rdb)

		ctx := context.Background()

		item := model.CreateArticle(
			article.URL("https://example.com"),
			article.Title("title"),
			article.Description("description"),
			article.Thumbnail("https://example.com"),
			article.TagList([]article.Tag{
				article.Tag("tag1"),
				article.Tag("tag2"),
			}),
		)

		if err := ag.Save(ctx, item); err != nil {
			t.Fatal(err)
		}

		got, err := ag.FindAll(ctx, repository.Index(2), repository.Size(2))
		if err != nil {
			t.Fatal(err)
		}

		articles := []model.Article{}

		if !reflect.DeepEqual(got, articles) {
			t.Errorf("FindAll() = %v, want %v", got, articles)
		}
	})
}
