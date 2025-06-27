package main

import (
	"math"
	"sort"
)

func superiorElements(a []int) []int {
	n := len(a)
	ans := []int{}
	maxi := math.MinInt

	for i := n - 1; i > 0; i-- {
		if a[i] > maxi {
			ans = append(ans, a[i])
		}
		if a[i] > maxi {
			maxi = a[i]
		}
	}
	sort.Ints(ans)
	return ans
}
