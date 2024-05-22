package main

import "fmt"

//func findKthLargest(nums []int, k int) int {
//	n := len(nums)
//	return quickselect(nums, 0, n-1, n-k)
//}

func quickselect(nums []int, l int, r int, k int) int {
	if l == r {
		return nums[k]
	}
	i := l - 1
	j := r + 1
	p := nums[l]
	for i < j {
		for i++; nums[i] < p; i++ {
		}
		for j--; nums[j] > p; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k < j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 2))
}
