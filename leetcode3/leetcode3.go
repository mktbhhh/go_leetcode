package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	max := 1
	from := 0
	to := 0
	hashMap := make(map[int32]int)
	for index, s := range s {
		p, ok := hashMap[s]
		if ok && p >= from {
			hashMap[s] = index
			from = p + 1
			to = index + 1
		} else {
			hashMap[s] = index
			to = index + 1
		}

		if to-from > max {
			max = to - from
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
}
