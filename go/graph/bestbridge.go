package graph

func BestBridge(grid [][]string) int {
	mainIsland := [][]int{}
	island := [][]int{}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			potentialIsland := traverseIsland(grid, row, col, &island)
			if len(potentialIsland) > 0 && len(mainIsland) == 0 {
				mainIsland = potentialIsland
			}
		}
	}


	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	queue := [][]int{}
	for _, pos := range mainIsland {
		queue = append(queue, []int{pos[0], pos[1], 0})
		visited[pos[0]][pos[1]] = true
	}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
	
		row, col, distance := pos[0], pos[1], pos[2] // destructure syntax
		neighbors := [][]int{{row + 1, col}, {row - 1, col}, {row, col + 1}, {row, col - 1}}
		for _, neighbor := range neighbors {
			nr, nc := neighbor[0], neighbor[1]
			if isInbounds(grid, nr, nc) && !visited[nr][nc] {
				if grid[nr][nc] == "L" {
					return distance
				}
				if grid[nr][nc] == "W" {
					queue = append(queue, []int{nr, nc, distance + 1})
					visited[nr][nc] = true
				}
			}
		}
	}
	return -1
}

func isInbounds(grid [][]string, row, col int) bool {
	inboundsRow := row >= 0 && row < len(grid)
	inboundsCol := col >= 0 && col < len(grid[0])
	return inboundsRow && inboundsCol
}

func contains(slice [][]int, item []int) bool {
	for _, i := range slice {
		if i[0] == item[0] && i[1] == item[1] {
			return true
		}
	}
	return false
}

func traverseIsland(grid [][]string, row int, col int, island *[][]int) [][]int {
 if !isInbounds(grid, row, col) || grid[row][col] == "W" {
 	return *island
 }
 if contains(*island, []int{row, col}) {
 	return *island
 }

 *island = append(*island, []int{row, col})
 traverseIsland(grid, row+1, col, island)
 traverseIsland(grid, row-1, col, island)
 traverseIsland(grid, row, col+1, island)
 traverseIsland(grid, row, col-1, island)
 return *island
}

