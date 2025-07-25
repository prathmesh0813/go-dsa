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
	// result := FindDuplicateValuesInArray([]int{3, 4, 5, 1, 2, 2})
	// fmt.Println(result)

	// //Problem-no-05-threesum problem
	// result := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	// fmt.Println(result)

	//Problem01 find second largest and second smallest
	// arr := []int{2, 4, 5, 7, 2}
	// n := len(arr)

	// fmt.Println("second largest is ", secondLagrest(arr, n))
	// fmt.Print("second smallest is ", seconSmallest(arr, n))

	//Problem02 Left rotate an Array by one
	// arr := []int{1, 2, 3, 4, 5}
	// n := len(arr)
	// fmt.Println("Rotated array:", rotateArray(arr, n))

	//Problem03 Right rotate an Array by one
	// arr := []int{1, 2, 3, 4, 5}
	// n := len(arr)
	// fmt.Println("Rotated array:", rotateArrayToRight(arr, n))

	//Problem04 Union of two sorted arrays
	// a := []int{1, 2, 2, 3, 4}
	// b := []int{2, 3, 5, 6, 10}

	// result := UnionOfArray(a, b)
	// fmt.Println("Union of arrays:", result)

	//Problem05 Alternate number
	//Problem Statement
	// There is an array 'A' of size 'N' with positive and negative elements. Without altering the relative order of positive and negative
	//  numbers, you must return an array of alternative positive and negative values.
	// Note:Start the array with a positive number. If any of the positive and negative numbers are left, add them at the end without
	// altering the order.
	// a := []int{3, 1, -2, -5, 2, -4, -9, 12, 13}
	// res := alternateNumbers(a)
	// fmt.Println(res)

	//Problem06 Superior Element or Leader element
	//There is an integer array 'A' of size 'N'. An element is called a Superior Element if it is greater than all the elements present
	//to its right.You must return a sorted array of all Superior Elements in the array "A.
	// Note:The last element of the array is always a Superior Element

	a := []int{1, 2, 3, 23, 4, 22, 21}
	result := superiorElements(a)
	fmt.Println(result)
}
