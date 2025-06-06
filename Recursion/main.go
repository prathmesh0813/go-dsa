package main

import "fmt"

func main() {
	//Backtracking01
	// n := 4
	// print1toN(1,n)

	//Backtracking02
	// arr := []int{1, 2, 1}
	// n := len(arr)
	// sum := 2
	//ds := []int {}

	//printSum(0,ds,0,sum,arr,n)

	// //Backtracking03
	// arr := []int{1, 2, 1}
	// n := len(arr)
	// sum := 2
	// result := printCount(0, 0, sum, arr, n)
	// fmt.Print(result)

	//Backtracking04
	// arr := [] int {3,1,2}
	// fmt.Println(subsetSum(arr))

	//Backtracking05
	// nums := []int{1, 2, 3}
    // result := permute(nums)
    // fmt.Println(result)

	//Backtracking06
	// G := [][]int{
	// 	{1, 2},  // Node 0 is connected to 1 and 2
	// 	{0, 2},  // Node 1 is connected to 0 and 2
	// 	{0, 1},  // Node 2 is connected to 0 and 1
	// }
	// m := 3 // Number of colors

	// if graphColoring(G, m) {
	// 	fmt.Println("Graph can be colored with", m, "colors.")
	// } else {
	// 	fmt.Println("Graph cannot be colored with", m, "colors.")
	// }

	//Backtracking07
	maze := [][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
	}

	n := 4
	paths := findPath(maze, n)
	fmt.Println("Possible paths:", paths)
}