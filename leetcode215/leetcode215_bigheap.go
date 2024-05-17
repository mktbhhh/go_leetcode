package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildHeap(nums, heapSize)
	for i := heapSize - 1; i >= len(nums)-k+1; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapSize--
		heapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildHeap(nums []int, heapSize int) {
	for i := heapSize/2 + 1; i >= 0; i-- {
		heapify(nums, i, heapSize)
	}
}

func heapify(nums []int, root int, heapSize int) {
	l, r, maxIndex := root*2+1, root*2+2, root
	if l < heapSize && nums[l] > nums[maxIndex] {
		maxIndex = l
	}
	if r < heapSize && nums[r] > nums[maxIndex] {
		maxIndex = r
	}

	if root != maxIndex {
		nums[maxIndex], nums[root] = nums[root], nums[maxIndex]
		heapify(nums, maxIndex, heapSize)
	}
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 4))
}
