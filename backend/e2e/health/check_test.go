//go:build e2e
// +build e2e

package health_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-guild/platform/e2e/health"
	"github.com/morning-night-guild/platform/e2e/helper"
	healthv1 "github.com/morning-night-guild/platform/pkg/connect/proto/health/v1"
)

func TestE2EHealthCheck(t *testing.T) {
	t.Parallel()

	url := helper.GetEndpoint(t)

	t.Run("ヘルスチェックが成功する", func(t *testing.T) {
		t.Parallel()

		client := health.New(t, http.DefaultClient, url)

		req := &healthv1.CheckRequest{}

		_, err := client.Health.Check(context.Background(), connect.NewRequest(req))
		if err != nil {
			t.Errorf("failed to health check: %s", err)
		}
	})
}
