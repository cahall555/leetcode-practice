package graph

func UndirectedPath(edges [][]string, nodeA string, nodeB string) bool {
	graph := buildGraph(edges)
	return hasPath(graph, nodeA, nodeB, map[string]bool{}) //map[string]bool{} is an empty set, when populated it will look something like {"A": true, "B": true}
}

func hasPath(graph map[string][]string, src string, dst string, visited map[string]bool) bool {
	if src == dst {
		return true
	}
	if visited[src] {
		return false
	}
	visited[src] = true

	for _, neighbor := range graph[src] {
		if hasPath(graph, neighbor, dst, visited) {
			return true
		}
	}

	return false
}

func buildGraph(edges [][]string) map[string][]string {
	graph := map[string][]string{}

	for _, edge := range edges {
		node1 := edge[0]
		node2 := edge[1]

		if _, found := graph[node1]; !found {
			graph[node1] = []string{}
		}
		if _, found := graph[node2]; !found {
			graph[node2] = []string{}
		}

		graph[node1] = append(graph[node1], node2)
		graph[node2] = append(graph[node2], node1)
	}
	return graph
}
