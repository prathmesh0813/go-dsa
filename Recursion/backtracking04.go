package main

import "sort"

//subset of sum in ascending order using backtracking
func subsetSum(candidates []int) []int {
	var res []int
	var backtrack func(ind int, sum int)
	backtrack = func(ind int, sum int) {
		if ind == len(candidates) {
			res = append(res, sum)
			return
		}
		//Include the element
		backtrack(ind+1, sum+candidates[ind])

		//Exclude the element
		backtrack(ind+1, sum)
	}

	backtrack(0, 0)
	sort.Ints(res)
	return res
}