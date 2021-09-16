package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}

	for xi, x := range nums {
		if yi, ok := hashTable[target-x]; ok {
			return []int{yi, xi}
		}
		hashTable[x] = xi
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
