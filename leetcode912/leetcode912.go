package main

import "fmt"

func sortArray(nums []int) []int {
	pos := 0
	nLen := len(nums)
	nums[pos], nums[nLen-1] = nums[nLen-1], nums[pos]
	quickSort(nums, 0, nLen-1)

	return nums
}

func quickSort(nums []int, l int, r int) {
	if l < r {
		pos := partition(nums, l, r)
		quickSort(nums, l, pos-1)
		quickSort(nums, pos+1, r)
	}
}

func partition(nums []int, l int, r int) int {
	pivot := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] < pivot {
			i += 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]

	return i + 1
}

func main() {
	nums := []int{0, 1, 1, 2, 0}
	sortArray(nums)
	fmt.Println(nums)
}
