package main

func LowerBound(nums []int, n, target int) int {
	low, high := 0, n-1
	ans := n
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] >= target {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}
