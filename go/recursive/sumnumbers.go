package recursive

// This function is to practice recursion in Go. It is not the most effient way to solve the problem. An iterative solution would be better. 
//O(n^2) time and space complexity.


func SumNumbers(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + SumNumbers(numbers[1:])
}
