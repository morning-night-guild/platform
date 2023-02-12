package helper

import (
	"net/http"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-guild/platform/pkg/connect/proto/article/v1/articlev1connect"
	"github.com/morning-night-guild/platform/pkg/connect/proto/health/v1/healthv1connect"
)

type ConnectClient struct {
	Article articlev1connect.ArticleServiceClient
	Health  healthv1connect.HealthServiceClient
}

func NewConnectClient(t *testing.T, client connect.HTTPClient, url string) *ConnectClient {
	t.Helper()

	ac := articlev1connect.NewArticleServiceClient(
		client,
		url,
	)

	hc := healthv1connect.NewHealthServiceClient(
		client,
		url,
	)

	return &ConnectClient{
		Article: ac,
		Health:  hc,
	}
}

type APIKeyTransport struct {
	t         *testing.T
	APIKey    string
	Transport http.RoundTripper
}

func NewAPIKeyTransport(
	t *testing.T,
	key string,
) *APIKeyTransport {
	t.Helper()

	return &APIKeyTransport{
		t:         t,
		APIKey:    key,
		Transport: http.DefaultTransport,
	}
}

func (at *APIKeyTransport) transport() http.RoundTripper {
	at.t.Helper()

	return at.Transport
}

func (at *APIKeyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	at.t.Helper()

	req.Header.Add("X-Api-Key", at.APIKey)

	resp, err := at.transport().RoundTrip(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
