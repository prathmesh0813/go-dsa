package main

import "fmt"

//print the sequences whose sum is K
func printSum(ind int, ds []int, s int, sum int, arr []int, n int) {
	if ind == n {
		if s == sum {
			for _, val := range ds {
				fmt.Print(val," ")
			}
			fmt.Println()
		}
		return
	}

	//Pick the element
	ds = append(ds, arr[ind])
	
	printSum(ind+1,ds,s+arr[ind],sum,arr,n)

	//Backtracking
	ds = ds[:len(ds)-1]
	printSum(ind+1,ds,s,sum,arr,n)
}