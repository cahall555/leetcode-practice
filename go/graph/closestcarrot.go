package graph


func ClosestCarrot(grid [][]string, startRow int, startCol int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	visited[startRow][startCol] = true
	queue := [][]int{{startRow, startCol, 0}}
	
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		
		row, col, distance := pos[0], pos[1], pos[2]
		neighbors := [][]int{{row + 1, col}, {row - 1, col}, {row, col + 1}, {row, col - 1}}
		for _, neighbor := range neighbors {
			nr, nc := neighbor[0], neighbor[1]
			if isInbounds(grid, nr, nc) && !visited[nr][nc] {
				if grid[nr][nc] == "C" {
					return distance + 1
				}
				if grid[nr][nc] != "X" {
					queue = append(queue, []int{nr, nc, distance + 1})
					visited[nr][nc] = true
				}
			}
		}
	}
	return -1
}


