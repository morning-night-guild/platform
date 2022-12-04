package controller_test

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-guild/platform/app/core/adapter/controller"
	healthv1 "github.com/morning-night-guild/platform/pkg/connect/proto/health/v1"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	os.Setenv("API_KEY", "test")

	type fields struct {
		apiKey string
	}

	type args struct {
		ctx context.Context
		req *connect.Request[healthv1.CheckRequest]
	}

	tests := []struct {
		name    string
		fields  fields
		h       controller.Health
		args    args
		want    *connect.Response[healthv1.CheckResponse]
		wantErr bool
	}{
		{
			name: "ヘルスチェックができる",
			fields: fields{
				apiKey: "test",
			},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[healthv1.CheckRequest]{},
			},
			want:    connect.NewResponse(&healthv1.CheckResponse{}),
			wantErr: false,
		},
		{
			name: "X-API-KEYが不正の時に認証エラーになる",
			fields: fields{
				apiKey: "invalid-api-key",
			},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[healthv1.CheckRequest]{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := controller.NewHealth()
			tt.args.req.Header().Set("X-API-KEY", tt.fields.apiKey)
			got, err := h.Check(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Health.Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Health.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
