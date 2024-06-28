package graph

func IslandCount(grid [][]string) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	count := 0

	for row := 0; row <= len(grid); row ++ {
		for col := 0; col <= len(grid[0]); col ++ {
			if exploreIsland(grid, row, col, &visited) {
				count ++
			}
		}
	}
	return count
}

func exploreIsland(grid [][]string, row int, col int, visited *[][]bool) bool {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return false
	}

	if grid[row][col] == "W" {
		return false
	}

	if (*visited)[row][col] == true {
		return false
	}

	(*visited)[row][col] = true

	exploreIsland(grid, row + 1, col, visited)
	exploreIsland(grid, row - 1, col, visited)
	exploreIsland(grid, row, col + 1, visited)
	exploreIsland(grid, row, col - 1, visited)

	return true
}
