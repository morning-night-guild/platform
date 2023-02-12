package helper

import (
	"testing"

	"github.com/morning-night-guild/platform/pkg/openapi"
)

type OpenAPIClient struct {
	Client *openapi.Client
}

func NewOpenAPIClient(t *testing.T, url string) *OpenAPIClient {
	t.Helper()

	client, err := openapi.NewClient(url + "/api")
	if err != nil {
		t.Error(err)
	}

	return &OpenAPIClient{
		Client: client,
	}
}
