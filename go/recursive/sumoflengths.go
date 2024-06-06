package recursive

// O(n^2) time/space complexity

func SumOfLengths(strings []string) int {
	if len(strings) == 0 {
		return 0
	}
	return len(strings[0]) + SumOfLengths(strings[1:])
}
