package helper

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func GenerateIDs(t *testing.T, count int) []uuid.UUID {
	t.Helper()

	ids := make([]uuid.UUID, count)

	for i := 0; i < count; i++ {
		ids[i] = uuid.MustParse(fmt.Sprintf("00000000-0000-0000-0000-%012d", i))
	}

	return ids
}
