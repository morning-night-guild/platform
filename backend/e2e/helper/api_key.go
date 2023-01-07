package helper

import (
	"os"
	"testing"
)

func GetAPIKey(t *testing.T) string {
	t.Helper()

	k := os.Getenv("API_KEY")

	if k == "" {
		k = "e2e"
	}

	return k
}
