package helper

import (
	"net/http"
	"testing"
)

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
