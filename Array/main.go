package main

import "fmt"

func main() {

	// //Reverse of array
	// ReverseInPlace([]int{1, 2, 3, 4, 5, 6}, 1, 4)

	//Add slice of two Nos.
	//AddSliceOfTwoNos([]int{9, 9, 9}, []int{1})

	//Sort a slice of unsorted integers using the Bubble Sort algorithm.
	result := BubbleSort([]int{2, 3, 1, 65, 5})
	fmt.Println(result)
}
