package main

func main() {
	//Backtracking01 
	// n := 4
	// print1toN(1,n)

	//Backtracking02

	arr := [] int {1,2,1}
	n := len(arr)
	sum := 2
	ds := []int {}

	printSum(0,ds,0,sum,arr,n)
}