package main

func rotateArrayToRight(arr []int, n int) []int {
	if n == 0 {
		return arr
	}
	temp := arr[n-1]
	for i := n - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = temp
	return arr
}
