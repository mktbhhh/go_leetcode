package main

import (
	"fmt"
	"math/rand"
)

func sortArray(nums []int) []int {
	qSort(nums, 0, len(nums)-1)
	return nums
}

func qSort(nums []int, left int, right int) {
	if left < right {
		pivot := rand.Int()%(right-left+1) + left
		nums[pivot], nums[right] = nums[right], nums[pivot]
		l := left
		r := right - 1
		pivot = nums[right]
		for l <= r {
			for ; l <= r && nums[l] < pivot; l++ {
			}
			for ; l <= r && nums[r] > pivot; r-- {
			}
			if l <= r {
				nums[l], nums[r] = nums[r], nums[l]
				l++
				r--
			}
		}
		nums[l], nums[right] = nums[right], nums[l]
		qSort(nums, left, l-1)
		qSort(nums, l+1, right)
	}
}

func main() {
	nums := []int{5, 2, 3, 1}
	sortArray(nums)
	fmt.Println(nums)
}
