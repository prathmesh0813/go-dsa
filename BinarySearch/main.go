package main

import "fmt"

func main() {

	//Problem01 = Implement an lower bound
	//You are given an array sorted in non decreasing order and a number 'target'.Find the index of lower bound of 'target'

	// nums := []int{1, 2, 3, 3, 5, 8, 8, 10, 10, 11}
	// target := 7
	// n := len(nums)
	// result := LowerBound(nums, n, target)
	// fmt.Println("Lower bound of", target, "having index:", result)

	// Problem02 - Implement an upper bound
	// You are given an array sorted in non-decreasing order and a number 'target'.
	// Find the index of the upper bound of 'target'.
	// Upper bound is defined as the first index where element > target

	// nums := []int{1, 2, 3, 3, 5, 8, 8, 10, 10, 11}
	// target := 11
	// n := len(nums)
	// result := UpperBound(nums, n, target)
	// fmt.Println("Upper bound of", target, "having index:", result)

	//Problem03
	// Floor() returns the greatest number in the array that is less than or equal to the target.
	// If no such number exists, it returns -1.
	// We assume the array is sorted in non-decreasing order.
	nums := []int{1, 2, 3, 3, 5, 8, 8, 10, 10, 11}
	target := 7
	n := len(nums)
	floorVal := Floor(nums, n, target)
	if floorVal != -1 {
		fmt.Println("Floor of", target, "is", floorVal)
	} else {
		fmt.Println("No floor found for", target)
	}
}
