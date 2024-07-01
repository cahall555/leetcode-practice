package dynamic

import (
	"testing"
)

func TestMinChange( t *testing.T) {
	coins := []int{1, 2, 5}
	if MinChange(11, coins) != 3 {
		t.Error("Expected 3")
	}

	coinsTwo := []int{1, 9, 5, 14, 30}
	if MinChange(13, coinsTwo) != 5 {
		t.Error("Expected 5")
	}
}
