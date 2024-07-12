package dynamic


func CanConcat(s string, words []string, memo map[string]bool) bool {
	if len(s) == 0 {
		return true
	}
	if val, ok := memo[s]; ok {
		return val
	}
	for _, word := range words {
		if len(word) <= len(s) && s[:len(word)] == word {
			if CanConcat(s[len(word):], words, memo) {
				memo[s] = true
				return true
			}
		}
	}
	memo[s] = false
	return false
}
