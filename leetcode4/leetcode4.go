package main

import "fmt"

/**
O(m+n)
*/
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)

	if len1 == 0 {
		if len2%2 == 0 {
			return float64((nums2[len2/2-1] + nums2[len2/2])) / 2
		} else {
			return float64(nums2[len2/2])
		}
	}

	if len2 == 0 {
		if len1%2 == 0 {
			return float64((nums1[len1/2-1] + nums1[len1/2])) / 2
		} else {
			return float64(nums1[len1/2])
		}
	}

	len := len1 + len2
	n1, n2 := 0, 0
	flag := false

	if len%2 == 0 {
		for i := 0; i < len/2; i++ {
			if n1 >= len1 {
				n2++
				flag = false
				continue
			}

			if n2 >= len2 {
				n1++
				flag = true
				continue
			}

			if nums1[n1] > nums2[n2] {
				n2++
				flag = false
			} else {
				n1++
				flag = true
			}
		}

		if flag {
			if n1 == len1 {
				return float64((nums1[n1-1] + nums2[n2])) / 2
			} else {
				if n2 == len2 {
					return float64((nums1[n1-1] + nums1[n1])) / 2
				} else {
					if nums1[n1] > nums2[n2] {
						return float64((nums1[n1-1] + nums2[n2])) / 2
					} else {
						return float64((nums1[n1-1] + nums1[n1])) / 2
					}
				}
			}

		} else {
			if n2 == len2 {
				return float64((nums2[n2-1] + nums1[n1])) / 2
			} else {
				if n1 == len1 {
					return float64((nums2[n2-1] + nums2[n2])) / 2
				} else {
					if nums1[n1] > nums2[n2] {
						return float64((nums2[n2-1] + nums2[n2])) / 2
					} else {
						return float64((nums2[n2-1] + nums1[n1])) / 2
					}
				}
			}
		}
	} else {
		for i := 0; i <= len/2; i++ {
			if n1 >= len1 {
				n2++
				flag = false
				continue
			}

			if n2 >= len2 {
				n1++
				flag = true
				continue
			}

			if nums1[n1] > nums2[n2] {
				n2++
				flag = false
			} else {
				n1++
				flag = true
			}
		}

		if flag {
			return float64(nums1[n1-1])
		} else {
			return float64(nums2[n2-1])
		}
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	nums1 := []int{1}
	nums2 := []int{2, 3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
