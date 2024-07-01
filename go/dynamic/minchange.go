package dynamic

import (
	"math"
)

func MinChange(amount int, coins []int) int {
	memo := make(map[int]int)
	answer := minChange(amount, coins, memo)
	if answer == math.MaxInt32 {
		return -1
	}
	return answer
}

func minChange(amount int, coins []int, memo map[int]int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return math.MaxInt32
	}

	if val, ok := memo[amount]; ok {
		return val
	}

	minCoins := math.MaxInt32
	for _, coin := range coins {
		remainder := amount - coin
		res := minChange(remainder, coins, memo)
		if res != math.MaxInt32 {
			minCoins = int(math.Min(float64(minCoins), float64(res+1)))
		}
	}

	memo[amount] = minCoins
	return minCoins
}
