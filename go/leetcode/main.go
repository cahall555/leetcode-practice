package main

import (
	"fmt"
	"example.com/dynamic"
	"example.com/recursive"
)

func main() {
	fmt.Println(dynamic.Tribonacci(14))

	fmt.Println(recursive.SumNumbers([]int{1, 2, 3, 4, 5}))
}
