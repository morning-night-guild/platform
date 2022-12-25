package env_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/morning-night-guild/platform/app/core/driver/env"
)

func TestInit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		env  string
		want env.Env
	}{
		{
			name: "prod",
			env:  "prod",
			want: "prod",
		},
		{
			name: "prev",
			env:  "prev",
			want: "prev",
		},
		{
			name: "dev",
			env:  "dev",
			want: "dev",
		},
		{
			name: "local",
			env:  "local",
			want: "local",
		},
		{
			name: "empty",
			env:  "",
			want: "",
		},
		{
			name: "other",
			env:  "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("ENV", "")
			env.Init()
			os.Setenv("ENV", tt.env)
			env.Init()
			if !reflect.DeepEqual(env.Get(), tt.want) {
				t.Errorf("Env = %v, want %v", env.Get(), tt.want)
			}
		})
	}
}
