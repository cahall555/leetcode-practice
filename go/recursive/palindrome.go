package recursive

// O(n^2) time/space complexity

func Palindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}

	return Palindrome(s[1: len(s)-1])
}
