package article_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/morning-night-guild/platform/app/core/model/article"
	interactor "github.com/morning-night-guild/platform/app/core/usecase/interactor/article"
	"github.com/morning-night-guild/platform/app/core/usecase/mock"
	"github.com/morning-night-guild/platform/app/core/usecase/port"
	"github.com/morning-night-guild/platform/app/core/usecase/repository"
)

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
				articleRepository: &mock.Article{
					Err: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.ShareArticleInput{
					URL:         article.URL("https://example.com"),
					Title:       article.Title("title"),
					Description: article.Description("description"),
					Thumbnail:   article.Thumbnail("https://example.com"),
				},
			},
			wantErr: false,
		},
		{
			name: "記事Repositoryのerrorを握りつぶさない",
			fields: fields{
				articleRepository: &mock.Article{
					Err: errors.New("article repository error"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.ShareArticleInput{
					URL:         article.URL("https://example.com"),
					Title:       article.Title("title"),
					Description: article.Description("description"),
					Thumbnail:   article.Thumbnail("https://example.com"),
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
