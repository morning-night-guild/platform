package config_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/morning-night-guild/platform/app/core/driver/config"
)

func TestInit(t *testing.T) { //nolint:tparallel
	t.Parallel()

	type args struct {
		dsn             string
		apiKey          string
		port            string
		newRelicAppName string
		newRelicLicense string
	}

	tests := []struct {
		name string
		args args
		want config.Config
	}{
		{
			name: "configを作成できる",
			args: args{
				dsn:             "dsn",
				apiKey:          "api_key",
				port:            "8080",
				newRelicAppName: "new_relic_app_name",
				newRelicLicense: "new_relic_license",
			},
			want: config.Config{
				DSN:             "dsn",
				APIKey:          "api_key",
				Port:            "8080",
				NewRelicAppName: "new_relic_app_name",
				NewRelicLicense: "new_relic_license",
			},
		},
		{
			name: "PORTの指定がなくてもconfigを作成できる",
			args: args{
				dsn:             "dsn",
				apiKey:          "api_key",
				port:            "",
				newRelicAppName: "new_relic_app_name",
				newRelicLicense: "new_relic_license",
			},
			want: config.Config{
				DSN:             "dsn",
				APIKey:          "api_key",
				Port:            "8080",
				NewRelicAppName: "new_relic_app_name",
				NewRelicLicense: "new_relic_license",
			},
		},
		{
			name: "PORTに数値に変換できない文字列が指定されてもconfigを作成できる",
			args: args{
				dsn:             "dsn",
				apiKey:          "api_key",
				port:            "port",
				newRelicAppName: "new_relic_app_name",
				newRelicLicense: "new_relic_license",
			},
			want: config.Config{
				DSN:             "dsn",
				APIKey:          "api_key",
				Port:            "8080",
				NewRelicAppName: "new_relic_app_name",
				NewRelicLicense: "new_relic_license",
			},
		},
	}

	// 環境変数を操作するため直列でテスト
	for _, tt := range tests { //nolint:paralleltest
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("DATABASE_URL", tt.args.dsn)
			os.Setenv("API_KEY", tt.args.apiKey)
			os.Setenv("NEWRELIC_APP_NAME", tt.args.newRelicAppName)
			os.Setenv("NEWRELIC_LICENSE", tt.args.newRelicLicense)
			config.Init()
			if !reflect.DeepEqual(config.Get(), tt.want) {
				t.Errorf("Config = %v, want %v", config.Get(), tt.want)
			}
		})
	}
}
