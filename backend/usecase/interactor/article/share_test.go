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

var _ repository.Article = (*ArticleMock)(nil)

// Article 記事リポジトリのモック.
type ArticleMock struct {
	Err error
}

// Save 記事を保存するモックメソッド.
func (a *ArticleMock) Save(ctx context.Context, article model.Article) error {
	return a.Err
}

func TestShareInteractorExecute(t *testing.T) {
	t.Parallel()

	type fields struct {
		articleRepository repository.Article
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
			},
			args: args{
				ctx: context.Background(),
				input: port.ShareArticleInput{
					Title:     article.Title("タイトル"),
					URL:       article.URL("https://example.com/image"),
					Thumbnail: article.Thumbnail("https://example.com/image"),
				},
			},
			wantErr: false,
		},
		{
			name: "Repositoryのerrorを握りつぶさない",
			fields: fields{
				articleRepository: &ArticleMock{
					Err: errors.New("repository error"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.ShareArticleInput{
					Title:     article.Title("タイトル"),
					URL:       article.URL("https://example.com/image"),
					Thumbnail: article.Thumbnail("https://example.com/image"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := interactor.NewShareInteractor(tt.fields.articleRepository)
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
