package main

import "fmt"

func main() {

	//Problem01 = Implement an lower bound
	//You are given an array sorted in non decreasing order and a number 'target'.Find the index of lower bound of 'target'

	nums := []int{1, 2, 3, 3, 5, 8, 8, 10, 10, 11}
	target := 7
	n := len(nums)
	result := LowerBound(nums, n, target)
	fmt.Println("Lower bound of", target, "having index:", result)
}
