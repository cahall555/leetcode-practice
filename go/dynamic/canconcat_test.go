package dynamic

import (
	"testing"
)

func TestCanConcat(t *testing.T) {
	words := []string{"one", "none", "is"}
	memoConcat := make(map[string]bool)
	if !CanConcat("noneisone", words, memoConcat) {
		t.Error("Expected true, got false")
	}
}
