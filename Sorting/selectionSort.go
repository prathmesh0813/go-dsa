package main

func selectionSort(arr []int, n int) {
	for i := 0; i <= n-2; i++ {
		minimum := i
		for j := i; j <= n-1; j++ {
			if arr[j] < arr[minimum] {
				minimum = j
			}
		}
		arr[minimum], arr[i] = arr[i], arr[minimum]
	}
}
