package graph

import (
	"testing"
)

func TestKnightAttack(t *testing.T) {
	tests := []struct {
		n        int
		kr       int
		kc       int
		pr       int
		pc       int
		expected int
	}{
		{8, 1, 1, 2, 2, 2},
		{8, 1, 1, 2, 3, 1},
		{8, 0, 3, 4, 2, 3},
		{8, 0, 3, 5, 2, 4},
		{24, 4, 7, 19, 20, 10},
		{100, 21, 10, 0, 0, 11},
		{3, 0, 0, 1, 2, 1},
		{3, 0, 0, 1, 1, -1},
	}

for _, tt := range tests {
	actual := KnightAttack(tt.n, tt.kr, tt.kc, tt.pr, tt.pc)
	if actual != tt.expected {
		t.Errorf("KnightAttack(%d, %d, %d, %d, %d): expected %d, actual %d", tt.n, tt.kr, tt.kc, tt.pr, tt.pc, tt.expected, actual)
	}
}
}

func BenchmarkKnightAttack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KnightAttack(8, 1, 1, 2, 2)
	}
}
