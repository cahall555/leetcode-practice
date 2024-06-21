package graph

// KnightAttack takes in a board size n, knight position (kr, kc), and pawn position (pr, pc) and returns the minimum number of moves the knight needs to make to capture the pawn. If the knight cannot capture the pawn, return -1.

// n represents the lenght of the board
//Time complexity: O(N^2)
//Space complexity: O(N^2)

func KnightAttack(n int, kr int, kc int, pr int, pc int) int {
	visited := make([][]int, n)//create a 2D slice of size n
	for i := range visited {// initialize inner slices
		visited[i] = make([]int, n)
	}
	queue := [][]int{{kr, kc}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current[0] == pr && current[1] == pc {
			return visited[current[0]][current[1]]
		}
		neighbors := knightPossitions(current[0], current[1], n)
		for _, neighbor := range neighbors {
			if visited[neighbor[0]][neighbor[1]] == 0 {
				queue = append(queue, neighbor)
				visited[neighbor[0]][neighbor[1]] = visited[current[0]][current[1]] + 1
			}
		}
	}
	return -1
}

func knightPossitions(r int, c int, n int) [][]int {
	possitions := [][]int{{r - 2, c - 1}, {r - 2, c + 1}, {r - 1, c - 2}, {r - 1, c + 2}, {r + 1, c - 2}, {r + 1, c + 2}, {r + 2, c - 1}, {r + 2, c + 1}}
	validPossitions := [][]int{}
	for _, pos := range possitions {
		if pos[0] >= 0 && pos[0] < n && pos[1] >= 0 && pos[1] < n {
			validPossitions = append(validPossitions, pos)
		}
	}
	return validPossitions
}

// KnightAttach(8, 1, 1, 2, 2) returns 2

// Iteration 1: knight is at (1, 1) and pawn is at (2, 2)
// queue = [(1, 1)]
// neighbors = [(3, 2), (2, 3)]
// updated queue = [(3, 2), (2, 3)]
// visited = {[(1, 1)=0, (3, 2)=1, (2, 3)=1]}

// Iteration 2: knight is at (3, 2) and pawn is at (2, 2)
// queue = [(2, 3)]
// neighbors = [(1, 1), (1, 3), (2, 4), (4, 4), (5, 1), (5, 3)]
// updated queue = [(2, 3), (1, 3), (2, 4), (4, 4), (5, 1), (5, 3)]
// visited = {[(1, 1)=0, (3, 2)=1, (2, 3)=1, (1, 3)=2, (2, 4)=2, (4, 4)=2, (5, 1)=2, (5, 3)=2]}

// Iteration 3: knight is at (2, 3) and pawn is at (2, 2) 
// queue = [(1, 3), (2, 4), (4, 4), (5, 1), (5, 3)]
// neighbors = [(0, 1), (0, 5), (1, 2), (1, 4), (3, 4), (3, 2), (4, 1), (4, 5)]
// updated queue = [(1, 3), (2, 4), (4, 4), (5, 1), (5, 3), (1, 2), (3, 4), (4, 1)]
// visited = {[(1, 1)=0, (3, 2)=1, (2, 3)=1, (1, 3)=2, (2, 4)=2, (4, 4)=2, (5, 1)=2, (5, 3)=2, (1, 2)=3, (3, 4)=3, (4, 1)=3]}

