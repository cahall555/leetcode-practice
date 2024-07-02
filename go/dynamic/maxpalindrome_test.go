package dynamic

import (
	"testing"
)

func TestMaxPalindrome(t *testing.T) {
	s := "chartreusepugvicefree"
	j := len(s) - 1
	memo := make(map[string]int)
	got := MaxPalindrome(s, 0, j, memo)
	want := 7
	if got != want {
		t.Errorf("MaxPalindrome(%q, 0, %d, memo) = %d; want %d", s, j, got, want)
	}
}
