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

	println("Max Palindrome")
	s := "chartreusepugvicefree"
	j := len(s) - 1
	memo := make(map[string]int)
	fmt.Println(dynamic.MaxPalindrome(s, 0, j, memo))
	
	println("Count Paths")
	grid := [][]string{
		{"O", "O", "X"},
		{"O", "O", "O"},
		{"O", "O", "O"},
	}

	memo = make(map[string]int)
	fmt.Println(dynamic.CountPaths(grid, 0, 0, memo))

	println("Array Stepper")
	nums := []int{2, 3, 2, 0, 0, 1}
	arrayMemo := make(map[int]int)
	fmt.Println(dynamic.ArrayStepper(nums, 0, arrayMemo))


	println("Can Concat")
	words := []string{"one", "none", "is"}
	memoConcat := make(map[string]bool)
	fmt.Println(dynamic.CanConcat("noneisone", words, memoConcat))

	println("overlaping sequence")
	s1 := "abcde"
	s2 := "ace"
	memoOverlap := make(map[string]int)
	fmt.Println(dynamic.OverlapSequence(s1, s2, 0, 0, memoOverlap))

}
