package main

import (
	"fmt"
	"example.com/dynamic"
	"example.com/recursive"
)

func main() {
	fmt.Println(dynamic.Tribonacci(14))

	fmt.Println(recursive.SumNumbers([]int{1, 2, 3, 4, 5}))

	fmt.Println(recursive.Factorial(5))

	fmt.Println(recursive.SumOfLengths([]string{"Luna", "Maci", "Teddy"}))

	fmt.Println(recursive.ReverseString("Hello"))

	fmt.Println(recursive.Palindrome("racecar"))

	fmt.Println(recursive.Fibonacci(6))
}
