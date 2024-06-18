package graph

//Article on graphs in Go: https://medium.com/@snassr/graphs-with-go-golang-part-i-3e0f9392c294

// Recursive depth first search to find if there is a path between two nodes in a graph
//time complexity: O(N)
//space complexity: O(N)

func HasPathRec(graph *map[string][]string, src string, dst string) bool { //graph is represented as an agjacency list
	if src == dst {
		return true
	}
	for _, neighbor := range (*graph)[src] {
		if HasPathRec(graph, neighbor, dst) {
			return true
		}
	}
	return false
}


// Iterative depth first search to find if there is a path between two nodes in a graph
//time complexity: O(N)
//space complexity: O(N)

func HasPathIt(graph *map[string][]string, src string, dst string) bool {
	if src == dst {
		return true
	}
	queue := []string{src}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] 
		if current == dst {
			return true
		}
		for _, neighbor := range (*graph)[current] {
			queue = append(queue, neighbor)
		}
	}
	return false
}
