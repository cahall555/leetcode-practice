package dynamic

import (
	"fmt"
)

//Time complexity: O(n^2)
//Space complexity: O(n^2

func MaxPalindrome(s string, i, j int, memo map[string]int) int {
	if i == j {
		return 1
	}

	if i > j {
		return 0
	}

	key := fmt.Sprintf("%d-%d", i, j)
	if val, ok := memo[key]; ok {
		return val
	}

	if s[i] == s[j] {
		memo[key] = 2 + MaxPalindrome(s, i+1, j-1, memo)
	} else {
		memo[key] = Max(MaxPalindrome(s, i+1, j, memo), MaxPalindrome(s, i, j-1, memo))
	}
	return memo[key]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// s := "chartreusepugvicefree"
// j := len(s) - 1
// memo := make(map[string]int)
// fmt.Println(MaxPalindrome(s, 0, j, memo))

// iteration 1:
// i = 0, j = 22
// key = "0-22"
// s[0] = c, s[22] = e
// s[0] != s[22]
// Max(MaxPalindrome(s, 1, 22, memo), MaxPalindrome(s, 0, 21, memo))
// memo["1-22"] = 2
//
// iteration 2:
// i = 1, j = 22
// key = "1-22"
// s[1] = h, s[22] = e
// s[1] != s[22]
// Max(MaxPalindrome(s, 2, 22, memo), MaxPalindrome(s, 1, 21, memo))
// memo["2-22"] = 2
//
// iteration 3:
// i = 2, j = 22
// key = "2-22"
// s[2] = a, s[22] = e
// s[2] != s[22]
// Max(MaxPalindrome(s, 3, 22, memo), MaxPalindrome(s, 2, 21, memo))
// memo["3-22"] = 2
//
// iteration 4:
// i = 3, j = 22
// key = "3-22"
// s[3] = r, s[22] = e
// s[3] != s[22]
// Max(MaxPalindrome(s, 4, 22, memo), MaxPalindrome(s, 3, 21, memo))
// memo["4-22"] = 2
//
// iteration 5:
// i = 4, j = 22
// key = "4-22"
// s[4] = t, s[22] = e
// s[4] != s[22]
// Max(MaxPalindrome(s, 5, 22, memo), MaxPalindrome(s, 4, 21, memo))
// memo["5-22"] = 2
