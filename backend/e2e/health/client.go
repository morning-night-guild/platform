package health

import (
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-guild/platform/pkg/connect/proto/health/v1/healthv1connect"
)

type Client struct {
	Health healthv1connect.HealthServiceClient
}

func New(t *testing.T, client connect.HTTPClient, url string) *Client {
	t.Helper()

	c := healthv1connect.NewHealthServiceClient(
		client,
		url,
	)

	return &Client{
		Health: c,
	}
}
