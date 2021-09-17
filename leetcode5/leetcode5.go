package main

import "fmt"

func longestPalindrome(s string) string {
	len := len(s)
	if len < 2 {
		return s
	}
	maxlen := 1
	f := 0
	t := 0

	dp := make([][]bool, len)

	for i := 0; i < len; i++ {
		dp[i] = make([]bool, len)
	}

	for i := 0; i < len; i++ {
		dp[i][i] = true
	}
	for L := 2; L <= len; L++ {
		for i := 0; i < len; i++ {
			j := i + L - 1

			if j >= len {
				break
			}

			if s[i] == s[j] {
				if j-i == 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}

			if dp[i][j] && L > maxlen {
				maxlen = L
				f = i
				t = j
			}
		}
	}

	return s[f : t+1]
}

func longestPalindrome1(s string) string {
	len := len(s)
	// if len < 2 {
	// 	return s
	// }

	end := -1
	start := 0

	for i := 0; i < len; i++ {
		start1, end1 := extendString(s, i, i)
		start2, end2 := extendString(s, i, i+1)

		if end1-start1 > end-start {
			start, end = start1, end1
		}

		if end2-start2 > end-start {
			start, end = start2, end2
		}
	}

	return s[start : end+1]
}

func extendString(s string, left int, right int) (int, int) {
	for ; right < len(s) && left >= 0 && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

func main() {
	fmt.Println(longestPalindrome1(""))
}
