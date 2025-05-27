package main

//print the count of subsequences whose sum is K 
func printCount(ind int, s int, sum int, arr []int, n int) int{
	if ind == n {
		if s == sum {
			return 1
		}else{

			return 0
		}
	}

	//Pick the element
	
	left := printCount(ind+1,s+arr[ind],sum,arr,n)

	//Backtracking
	right := printCount(ind+1,s,sum,arr,n)

	return left + right
}