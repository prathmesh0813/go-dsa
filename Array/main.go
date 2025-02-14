package main

import "fmt"

func main() {

	// //Problem-no-01-Reverse of array
	// ReverseInPlace([]int{1, 2, 3, 4, 5, 6}, 1, 4)

	//Problem-no-02-Add slice of two Nos.
	//AddSliceOfTwoNos([]int{9, 9, 9}, []int{1})

	//Problem-no-03-Sort a slice of unsorted integers using the Bubble Sort algorithm.
	// result := BubbleSort([]int{2, 3, 1, 65, 5})
	// fmt.Println(result)

	//Problem-no-04-Given an unsorted slice of n positive integers like {3,2,1,4,5,4,…,n} where each number is smaller than n and there is at most one duplicate, return the duplicate value like 4.
	result := FindDuplicateValuesInArray([]int{3, 4, 5, 1, 2, 2})
	fmt.Println(result)
}
