package main

import "fmt"

func main() {
	//Backtracking01
	// n := 4
	// print1toN(1,n)

	//Backtracking02
	// arr := []int{1, 2, 1}
	// n := len(arr)
	// sum := 2
	//ds := []int {}

	//printSum(0,ds,0,sum,arr,n)

	// //Backtracking03
	// arr := []int{1, 2, 1}
	// n := len(arr)
	// sum := 2
	// result := printCount(0, 0, sum, arr, n)
	// fmt.Print(result)

	//Backtracking04
	// arr := [] int {3,1,2}
	// fmt.Println(subsetSum(arr))

	//Backtracking05
	nums := []int{1, 2, 3}
    result := permute(nums)
    fmt.Println(result)
}