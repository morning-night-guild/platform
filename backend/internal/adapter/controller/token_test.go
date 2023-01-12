package controller_test

import (
	"reflect"
	"testing"

	"github.com/morning-night-guild/platform/internal/adapter/controller"
	"github.com/morning-night-guild/platform/internal/usecase/repository"
)

func TestTokenToIndex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		tr   controller.Token
		want repository.Index
	}{
		{
			name: "トークンからインデックスを生成できる",
			tr:   controller.CreateTokenFromIndex(repository.Index(0)),
			want: repository.Index(0),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.tr.ToIndex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Token.ToIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenCreateNextToken(t *testing.T) {
	t.Parallel()

	type args struct {
		size repository.Size
	}

	tests := []struct {
		name string
		tr   controller.Token
		args args
		want controller.Token
	}{
		{
			name: "ネクストトークンが作成できる",
			tr:   controller.CreateTokenFromIndex(repository.Index(0)),
			args: args{
				size: repository.Size(1),
			},
			want: controller.CreateTokenFromIndex(repository.Index(1)),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.tr.CreateNextToken(tt.args.size); got != tt.want {
				t.Errorf("Token.CreateNextToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
