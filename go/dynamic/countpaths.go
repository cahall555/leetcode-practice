package dynamic

import (
	"fmt"
)

func CountPaths(grid [][]string, row int, col int, memo map[string]int) int {
	if row == len(grid) - 1 && col == len(grid[0]) - 1{
		return 1
	}

	if row == len(grid) || col == len(grid[0]) || grid[row][col] == "X"{
		return 0
	}

	key := fmt.Sprintf("%d,%d", row, col)
	if _, ok := memo[key]; ok {
		return memo[key]
	}

	memo[key] = CountPaths(grid, row + 1, col, memo) + CountPaths(grid, row, col + 1, memo)
	return memo[key]
}
