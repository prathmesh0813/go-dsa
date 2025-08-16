package main

import (
	"math"
)

func findKthElement(nums1 []int, nums2 []int, k int) int {
	n1, n2 := len(nums1), len(nums2)

	if n1 > n2 {
		return findKthElement(nums2, nums1, k)
	}

	low := int(math.Max(0, float64(k-n2)))
	high := int(math.Min(float64(k), float64(n1)))

	for low <= high {
		mid1 := (low + high) / 2
		mid2 := k - mid1

		l1, l2 := math.MinInt32, math.MinInt32
		r1, r2 := math.MaxInt32, math.MaxInt32

		if mid1 < n1 {
			r1 = nums1[mid1]
		}
		if mid2 < n2 {
			r2 = nums2[mid2]
		}
		if mid1-1 >= 0 {
			l1 = nums1[mid1-1]
		}
		if mid2-1 >= 0 {
			l2 = nums2[mid2-1]
		}

		if l1 <= r2 && l2 <= r1 {
			return max(l1, l2)
		} else if l1 > r2 {
			high = mid1 - 1
		} else {
			low = mid1 + 1
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
