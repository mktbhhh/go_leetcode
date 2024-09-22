package main

import "fmt"

func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}

	length := len(s)
	begin := 0
	maxLen := 1

	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true
	}

	for L := 2; L <= length; L++ {
		for i := 0; i <= length-L; i++ {
			j := i + L - 1

			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}

				if dp[i][j] && L > maxLen {
					maxLen = L
					begin = i
				}
			} else {
				dp[i][j] = false
			}
		}
	}

	return s[begin : begin+maxLen]
}

func main() {
	s := "bb"
	fmt.Println(longestPalindrome2(s))
}
