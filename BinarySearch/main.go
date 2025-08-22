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
	// nums := []int{1, 2, 3, 3, 5, 8, 8, 10, 10, 11}
	// target := 7
	// n := len(nums)
	// floorVal := Floor(nums, n, target)
	// if floorVal != -1 {
	// 	fmt.Println("Floor of", target, "is", floorVal)
	// } else {
	// 	fmt.Println("No floor found for", target)
	// }

	//Problem04 Find the nth root of integer using binary search
	// 3th root of integer 27 is 3
	// fmt.Println(NthRoot(3, 27))  // 3
	// fmt.Println(NthRoot(6, 64))  // 2
	// fmt.Println(NthRoot(3, 20))  // -1
	// fmt.Println(NthRoot(4, 1))   // 1
	// fmt.Println(NthRoot(4, 0))   // 0
	// fmt.Println(NthRoot(5, 243)) // 3

	// //Problem02 kth element in two sorted arrays using binary search
	// nums1 := []int{2, 3, 6, 7, 9}
	// nums2 := []int{1, 4, 8, 10}
	// k := 5

	// fmt.Println("Kth element is:", findKthElement(nums1, nums2, k))
	//Problem06 Median of row wise sorted matrix
	matrix := [][]int{
		{1, 3, 5},
		{2, 6, 9},
		{3, 6, 9},
	}
	fmt.Println(matrixMedian(matrix))
}
