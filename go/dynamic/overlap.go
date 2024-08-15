package dynamic

import (
	"fmt"
)

func OverlapSequence(s1, s2 string, i, j int, memo map[string]int) int {
	if i == len(s1) || j == len(s2) {
		return 0
	}

	key := fmt.Sprintf("%d-%d", i, j)
	if val, ok := memo[key]; ok {
		return val
	}

	if s1[i] == s2[j] {
		memo[key] = 1 + OverlapSequence(s1, s2, i+1, j+1, memo)
	} else {
		memo[key] = Max(OverlapSequence(s1, s2, i+1, j, memo), OverlapSequence(s1, s2, i, j+1, memo))
	}
	return memo[key]
}
