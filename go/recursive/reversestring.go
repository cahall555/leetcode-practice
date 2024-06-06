package recursive

// O(n^2) time/space complexity

func ReverseString(s string) string {
	if len(s) == 0 {
		return ""
	}

	return ReverseString(s[1:]) + string(s[0])
}
