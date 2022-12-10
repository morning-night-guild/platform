package controller_test

import (
	"context"
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/morning-night-guild/platform/app/core/adapter/controller"
	"github.com/morning-night-guild/platform/app/core/model"
	"github.com/morning-night-guild/platform/app/core/model/article"
	me "github.com/morning-night-guild/platform/app/core/model/errors"
	"github.com/morning-night-guild/platform/app/core/usecase"
	"github.com/morning-night-guild/platform/app/core/usecase/port"
	articlev1 "github.com/morning-night-guild/platform/pkg/connect/article/v1"
)

type ShareMock struct {
	Err error
}

func (s ShareMock) Execute(ctx context.Context, input port.ShareArticleInput) (port.ShareArticleOutput, error) {
	return port.ShareArticleOutput{}, s.Err
}

type ListMock struct {
	Articles []model.Article
	Err      error
}

func (l ListMock) Execute(ctx context.Context, input port.ListArticleInput) (port.ListArticleOutput, error) {
	return port.ListArticleOutput{
		Articles: l.Articles,
	}, l.Err
}

func TestArticleShare(t *testing.T) {
	t.Parallel()

	os.Setenv("API_KEY", "test")

	type fields struct {
		apiKey string
		share  usecase.Usecase[port.ShareArticleInput, port.ShareArticleOutput]
		list   usecase.Usecase[port.ListArticleInput, port.ListArticleOutput]
	}

	type args struct {
		ctx context.Context
		req *connect.Request[articlev1.ShareRequest]
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *connect.Response[articlev1.ShareResponse]
		wantErr error
	}{
		{
			name: "記事の共有ができる",
			fields: fields{
				apiKey: "test",
				share: ShareMock{
					Err: nil,
				},
				list: ListMock{},
			},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[articlev1.ShareRequest]{
					Msg: &articlev1.ShareRequest{
						Url: "https://example.com",
					},
				},
			},
			want:    connect.NewResponse(&articlev1.ShareResponse{}),
			wantErr: nil,
		},
		{
			name: "X-Api-Keyが不正の時、認証エラーになる",
			fields: fields{
				apiKey: "invalid-api-key",
				share: ShareMock{
					Err: nil,
				},
				list: ListMock{},
			},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[articlev1.ShareRequest]{
					Msg: &articlev1.ShareRequest{
						Url: "https://example.com",
					},
				},
			},
			want:    nil,
			wantErr: controller.ErrUnauthorized,
		},
		{
			name: "URLが不正の時、バッドリクエストエラーになる",
			fields: fields{
				apiKey: "test",
				share: ShareMock{
					Err: nil,
				},
				list: ListMock{},
			},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[articlev1.ShareRequest]{
					Msg: &articlev1.ShareRequest{
						Url: "http://example.com",
					},
				},
			},
			want:    nil,
			wantErr: controller.ErrInvalidArgument,
		},
		{
			name: "ユースケースでバリデーションエラーが発生した際、バッドリクエストエラーになる",
			fields: fields{
				apiKey: "test",
				share: ShareMock{
					Err: me.NewValidationError("validation error"),
				},
				list: ListMock{},
			},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[articlev1.ShareRequest]{
					Msg: &articlev1.ShareRequest{
						Url: "https://example.com",
					},
				},
			},
			want:    nil,
			wantErr: controller.ErrInvalidArgument,
		},
		{
			name: "ユースケースでバリデーションエラー以外のエラーが発生した際、サーバーエラーになる",
			fields: fields{
				apiKey: "test",
				share: ShareMock{
					Err: errors.New("unknown error"),
				},
			},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[articlev1.ShareRequest]{
					Msg: &articlev1.ShareRequest{
						Url: "https://example.com",
					},
				},
			},
			want:    nil,
			wantErr: controller.ErrInternal,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := controller.NewArticle(tt.fields.share, tt.fields.list)
			tt.args.req.Header().Set("X-Api-Key", tt.fields.apiKey)
			got, err := a.Share(tt.args.ctx, tt.args.req)
			if err != nil && err != tt.wantErr {
				t.Errorf("Article.Share() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Article.Share() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleList(t *testing.T) {
	t.Parallel()

	type fields struct {
		share usecase.Usecase[port.ShareArticleInput, port.ShareArticleOutput]
		list  usecase.Usecase[port.ListArticleInput, port.ListArticleOutput]
	}

	type args struct {
		ctx context.Context
		req *connect.Request[articlev1.ListRequest]
	}

	id := uuid.New()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *connect.Response[articlev1.ListResponse]
		wantErr bool
	}{
		{
			name: "記事の一覧が取得できる（ネクストトークンあり）",
			fields: fields{
				share: ShareMock{},
				list: ListMock{
					Articles: []model.Article{
						{
							ID:          article.ID(id),
							Title:       article.Title("title"),
							URL:         article.URL("https://example.com"),
							Description: article.Description("description"),
							Thumbnail:   article.Thumbnail("https://example.com"),
							TagList:     article.TagList{},
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[articlev1.ListRequest]{
					Msg: &articlev1.ListRequest{
						PageToken:   "",
						MaxPageSize: 1,
					},
				},
			},
			want: connect.NewResponse(&articlev1.ListResponse{
				Articles: []*articlev1.Article{
					{
						Id:          id.String(),
						Title:       "title",
						Url:         "https://example.com",
						Description: "description",
						Thumbnail:   "https://example.com",
						Tags:        []string{},
					},
				},
				NextPageToken: "MQ==",
			}),
			wantErr: false,
		},
		{
			name: "記事の一覧が取得できる（ネクストトークンなし）",
			fields: fields{
				share: ShareMock{},
				list: ListMock{
					Articles: []model.Article{
						{
							ID:          article.ID(id),
							Title:       article.Title("title"),
							URL:         article.URL("https://example.com"),
							Description: article.Description("description"),
							Thumbnail:   article.Thumbnail("https://example.com"),
							TagList:     article.TagList{},
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[articlev1.ListRequest]{
					Msg: &articlev1.ListRequest{
						PageToken:   "",
						MaxPageSize: 3,
					},
				},
			},
			want: connect.NewResponse(&articlev1.ListResponse{
				Articles: []*articlev1.Article{
					{
						Id:          id.String(),
						Title:       "title",
						Url:         "https://example.com",
						Description: "description",
						Thumbnail:   "https://example.com",
						Tags:        []string{},
					},
				},
				NextPageToken: "",
			}),
			wantErr: false,
		},
		{
			name: "不正なサイズを指定して記事の一覧が取得できない",
			fields: fields{
				share: ShareMock{},
				list:  ListMock{},
			},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[articlev1.ListRequest]{
					Msg: &articlev1.ListRequest{
						PageToken:   "",
						MaxPageSize: 1000,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := controller.NewArticle(tt.fields.share, tt.fields.list)
			got, err := a.List(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Article.List() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Article.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
