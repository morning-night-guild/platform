package server_test

import (
	"reflect"
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

func TestConvertAllowOrigins(t *testing.T) {
	type args struct {
		allowOrigins string
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "CORS許可のオリジンのリストを作成できる(単数)",
			args: args{
				allowOrigins: "http://example.com",
			},
			want:    []string{"http://example.com"},
			wantErr: false,
		},
		{
			name: "CORS許可のオリジンのリストを作成できる(複数)",
			args: args{
				allowOrigins: "http://example.com,http://example.com",
			},
			want:    []string{"http://example.com", "http://example.com"},
			wantErr: false,
		},
		{
			name: ",以外の文字で区切られた文字列の場合、1つのオリジンとして作成される",
			args: args{
				allowOrigins: "http://example.com http://example.com",
			},
			want:    []string{"http://example.com http://example.com"},
			wantErr: false,
		},
		{
			name: "空文字列の場合CORS許可のオリジンのリストを作成できない",
			args: args{
				allowOrigins: "",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := server.ConvertAllowOrigins(tt.args.allowOrigins)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertAllowOrigins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertAllowOrigins() = %v, want %v", got, tt.want)
			}
		})
	}
}
