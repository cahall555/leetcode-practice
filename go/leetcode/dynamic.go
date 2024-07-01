package main

import (
	"fmt"
	"os"
	"example.com/dynamic"
)

func main() {
	
	pageSize := os.Getpagesize()
	fmt.Printf("Page size: %d bytes (%d KB)\n", pageSize, pageSize/1024)
	
	println("Min Change")

	coins := []int{1, 2, 5}
	fmt.Println(dynamic.MinChange(11, coins))
	
	println("Tribonacci")
	fmt.Println(dynamic.Tribonacci(25))

}
