package main

import (
	"fmt"
	"os"
	"example.com/graph"
)

func main() {

	pageSize := os.Getpagesize()
    	fmt.Printf("Page size: %d bytes (%d KB)\n", pageSize, pageSize/1024)

var graphMap = map[string][]string{ //Acylic ajacency list
		"a": []string{"b", "c"},  // a -> b -> d -> f -> h
		"b": []string{"d"},        //|
		"c": []string{"e"},        //v
		"d": []string{"f"},        //c -> e -> g
		"e": []string{"g"},
		"f": []string{"h"},
	}

//	fmt.Println(graph.HasPathRec(&graphMap, "a", "h"))

	fmt.Println(graph.HasPathIt(&graphMap, "a", "h"))


	var edges = [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
	}

	fmt.Println(graph.UndirectedPath(edges, "i", "l"))


	var connected = map[int][]int{
		0: {8, 1, 5},
		1: {0},
		5: {0},
		8: {0},
		2: {3, 4},
		3: {2, 4},
		4: {3, 2},
	}

	fmt.Println(graph.ConnectedComponentCount(&connected))

	fmt.Println(graph.ConnectedComponentCountSet(&connected))

	fmt.Println(graph.KnightAttack(8, 1, 1, 2, 2))

	println("Best Bridge")
	var grid = [][]string{
		{"W", "W", "W", "L", "L"},
		{"L", "L", "W", "W", "L"},
		{"W", "L", "W", "W", "W"},
		{"W", "W", "W", "W", "W"},
		{"W", "W", "W", "W", "W"},
	}

	fmt.Println(graph.BestBridge(grid))

}
