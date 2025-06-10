package main

func rotateArray(arr []int, n int) []int {
	temp := arr[0]
	for i := 1; i < n; i++ {
		arr[i-1] = arr[i]
	}
	arr[n-1] = temp
	return arr
}
