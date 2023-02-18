package controller_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-guild/platform/internal/adapter/controller"
	me "github.com/morning-night-guild/platform/internal/domain/model/errors"
)

func TestControllerHandleConnectError(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		err error
	}

	tests := []struct {
		name     string
		args     args
		wantCode connect.Code
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				err: me.NewURLError(""),
			},
			wantCode: connect.CodeInvalidArgument,
		},
		{
			name: "",
			args: args{
				ctx: context.Background(),
				err: me.NewValidationError(""),
			},
			wantCode: connect.CodeInvalidArgument,
		},
		{
			name: "",
			args: args{
				ctx: context.Background(),
				err: errors.New("unknown error"),
			},
			wantCode: connect.CodeInternal,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctl := controller.New()
			err := ctl.HandleConnectError(tt.args.ctx, tt.args.err)
			if connectErr := new(connect.Error); errors.As(err, &connectErr) {
				code := connect.CodeOf(connectErr)
				if code != tt.wantCode {
					t.Errorf("HandleConnectError() = %v, want %v", code, tt.wantCode)
				}
			}
		})
	}
}
