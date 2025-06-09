package main

import "math"

func secondLagrest(arr []int, n int) int {
	largest := arr[0]
	slargest := -1

	for i := 1; i < n; i++ {
		if arr[i] > largest {
			slargest = largest
			largest = arr[i]
		} else if arr[i] < largest && arr[i] > slargest {
			slargest = arr[i]
		}
	}
	return slargest
}

func seconSmallest(arr []int, n int) int {
	smallest := arr[0]
	ssmallest := math.MaxInt

	for i := 1; i < n; i++ {
		if arr[i] < smallest {
			ssmallest = smallest
			smallest = arr[i]
		} else if arr[i] != smallest && arr[i] < ssmallest {
			ssmallest = arr[i]
		}
	}
	return ssmallest
}
