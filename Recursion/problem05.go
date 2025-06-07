package main

//Permutations using swapping(Optimize way)
func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func(ind int, path []int)
	backtrack = func(ind int, path []int) {
		if ind == len(nums){
				tmp := make([]int,len(path))
				copy(tmp,path)
				res = append(res, tmp)
				return
		}
		for i := ind; i < len(nums); i++ {
        swap(nums, i, ind)
        backtrack(ind+1, nums)
        swap(nums, i, ind) // backtrack
    }
	}
	backtrack(0,[]int {})
	return res
}

func swap(nums []int, i, j int) {
    nums[i], nums[j] = nums[j], nums[i]
}