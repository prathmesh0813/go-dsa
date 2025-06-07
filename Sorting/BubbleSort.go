package main

func bubbleSort(arr []int, n int) {
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
