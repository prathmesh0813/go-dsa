package main

func insertionSort(arr []int, n int) {
	for i := 0; i <= n-1; i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}

	}
}
