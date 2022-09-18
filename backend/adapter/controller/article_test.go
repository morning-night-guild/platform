package controller_test

import (
	"context"
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-guild/platform/adapter/controller"
	articlev1 "github.com/morning-night-guild/platform/driver/connect/article/v1"
	me "github.com/morning-night-guild/platform/model/errors"
	"github.com/morning-night-guild/platform/usecase"
	"github.com/morning-night-guild/platform/usecase/port"
)

type ShareMock struct {
	err error
}

func NewShareMock(err error) ShareMock {
	return ShareMock{
		err: err,
	}
}

func (s ShareMock) Execute(ctx context.Context, input port.ShareArticleInput) (port.ShareArticleOutput, error) {
	return port.ShareArticleOutput{}, s.err
}

func TestArticleShare(t *testing.T) {
	t.Parallel()

	os.Setenv("API_KEY", "test")

	type fields struct {
		apiKey string
		share  usecase.Usecase[port.ShareArticleInput, port.ShareArticleOutput]
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
				share:  NewShareMock(nil),
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
			name: "X-API-KEYが不正の時、認証エラーになる",
			fields: fields{
				apiKey: "invalid-api-key",
				share:  NewShareMock(nil),
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
				share:  NewShareMock(nil),
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
				share: NewShareMock(
					me.NewValidationError("validation error"),
				),
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
				share: NewShareMock(
					errors.New("unknown error"),
				),
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
			a := controller.NewArticle(tt.fields.share)
			tt.args.req.Header().Set("X-API-KEY", tt.fields.apiKey)
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
