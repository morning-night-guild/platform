package article_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/morning-night-guild/platform/model"
	"github.com/morning-night-guild/platform/model/article"
	interactor "github.com/morning-night-guild/platform/usecase/interactor/article"
	"github.com/morning-night-guild/platform/usecase/port"
	"github.com/morning-night-guild/platform/usecase/repository"
)

var (
	_ repository.Article = (*ArticleMock)(nil)
	_ repository.OGP     = (*OGPMock)(nil)
)

// Article 記事リポジトリのモック.
type ArticleMock struct {
	Err error
}

// Save 記事を保存するモックメソッド.
func (a *ArticleMock) Save(ctx context.Context, article model.Article) error {
	return a.Err
}

// OGP OGPリポジトリのモック.
type OGPMock struct {
	Article model.Article
	Err     error
}

// Create 記事を作成するモックメソッド.
func (o *OGPMock) Create(ctx context.Context, url article.URL) (model.Article, error) {
	if o.Err != nil {
		return model.Article{}, o.Err
	}

	o.Article.URL = url

	return o.Article, nil
}

func TestShareInteractorExecute(t *testing.T) {
	t.Parallel()

	type fields struct {
		articleRepository repository.Article
		ogpRepository     repository.OGP
	}

	type args struct {
		ctx   context.Context
		input port.ShareArticleInput
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    port.ShareArticleOutput
		wantErr bool
	}{
		{
			name: "記事を共有できる",
			fields: fields{
				articleRepository: &ArticleMock{
					Err: nil,
				},
				ogpRepository: &OGPMock{
					Article: model.CreateArticle(
						article.Title("タイトル"),
						article.URL("https://example.com"),
						article.Thumbnail("https://example.com/image"),
						article.TagList{},
					),
					Err: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.ShareArticleInput{
					URL: article.URL("https://example.com"),
				},
			},
			wantErr: false,
		},
		{
			name: "記事Repositoryのerrorを握りつぶさない",
			fields: fields{
				articleRepository: &ArticleMock{
					Err: errors.New("article repository error"),
				},
				ogpRepository: &OGPMock{
					Article: model.CreateArticle(
						article.Title("タイトル"),
						article.URL("https://example.com"),
						article.Thumbnail("https://example.com/image"),
						article.TagList{},
					),
					Err: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.ShareArticleInput{
					URL: article.URL("https://example.com"),
				},
			},
			wantErr: true,
		},
		{
			name: "OGPRepositoryのerrorを握りつぶさない",
			fields: fields{
				articleRepository: &ArticleMock{
					Err: nil,
				},
				ogpRepository: &OGPMock{
					Article: model.Article{},
					Err:     errors.New("ogp repository error"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.ShareArticleInput{
					URL: article.URL("https://example.com"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := interactor.NewShareInteractor(tt.fields.articleRepository, tt.fields.ogpRepository)
			got, err := s.Execute(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShareInteractor.Execute() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShareInteractor.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
