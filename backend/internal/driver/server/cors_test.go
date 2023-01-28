package server_test

import (
	"testing"

	"github.com/morning-night-guild/platform/internal/driver/server"
)

func TestNewCORS(t *testing.T) {
	t.Parallel()

	type args struct {
		allowOrigins []string
		debug        bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			// CORS設定の作成はライブラリを使用しているため、ここでは厳密にテストしない。
			name: "CORSの設定が作成できる",
			args: args{
				allowOrigins: []string{"http://example.com"},
				debug:        false,
			},
			wantErr: false,
		},
		{
			name: "許可するオリジンが空の場合、CORSの設定を作成できない",
			args: args{
				allowOrigins: []string{},
				debug:        false,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if _, err := server.NewCORS(tt.args.allowOrigins, tt.args.debug); (err != nil) != tt.wantErr {
				t.Errorf("NewCORS() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
		})
	}
}
