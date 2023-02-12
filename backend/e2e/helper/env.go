package helper

import (
	"os"
	"testing"
)

func GetAppCoreEndpoint(t *testing.T) string {
	t.Helper()

	return os.Getenv("APP_CORE_ENDPOINT")
}

func GetAPIKey(t *testing.T) string {
	t.Helper()

	return os.Getenv("API_KEY")
}
