package graph

// ConnectedComponentCount returns the number of connected components in a graph
// Time Complexity: O(n)
// Space Complexity: O(n)

// The two solutions are essintially the same but experament with two different ways of creating the visited variable.
// make(map[int]bool, len(*graph)) creates a map with a preallocated size of the length of the graph. This can be useful for optimization of larger projects.
// map[int]bool{} creates an empty map.
// The data structure for both looks something like: visited[0] = true, visited[1] = true ... visited[n] = true

func ConnectedComponentCount(graph *map[int][]int) int {
	visited := make(map[int]bool, len(*graph))
	count := 0
	for node := range *graph {
		if !visited[node] {
			depthFirstSearch(graph, node, visited)
			count++
		}
	}
	return count
}

func depthFirstSearch(graph *map[int][]int, node int, visited map[int]bool) {
	if visited[node] {
		return
	}
	visited[node] = true
	for _, neighbor := range (*graph)[node] {
		depthFirstSearch(graph, neighbor, visited)
	}
//	println("memory: ", visited)
}


func ConnectedComponentCountSet(graph *map[int][]int) int {
	visited := map[int]bool{}
	count := 0
	for node := range *graph {
		if explore(graph, node, visited) {
			count++
		}
	}
	return count
}

func explore(graph *map[int][]int, node int, visited map[int]bool) bool {
	if visited[node] {
		return false
	}
	visited[node] = true
	for _, neighbor := range (*graph)[node] {
		explore(graph, neighbor, visited)
	}
//	println("memory: ", visited)
	return true
}
