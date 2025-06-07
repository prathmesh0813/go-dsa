package main

import "fmt"

func main() {
	arr := []int{13, 46, 24, 52, 20, 9}
	n := len(arr)
	selectionSort(arr, n)
	fmt.Println(arr)
}
