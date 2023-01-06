package helper

import (
	"os"
	"testing"
)

func GetEndpoint(t *testing.T) string {
	t.Helper()

	ep := os.Getenv("ENDPOINT")

	if ep == "" {
		ep = "http://localhost:8081"
	}

	return ep
}
