package dynamic

import (
	"testing"
)

func TestOverlapSequence(t *testing.T) {
	s1 := "kinfolklivemustache"
	s2 := "bespokekinfolksnackwave"
	memoOverlap := make(map[string]int)
	result := OverlapSequence(s1, s2, 0, 0, memoOverlap)
	if result != 11 {
		t.Errorf("Expected 3 but got %d", result)
	}
}
