func subsetsWithDup(nums []int) [][]int {
    var res [][] int

    var backtrack func(ind int, path []int)
    backtrack = func(ind int,path []int){
        tmp := make([]int,len(path))
        copy(tmp,path)
        res = append(res,tmp)

        for i := ind;i<len(nums);i++ {
            if i != ind && nums[i] == nums[i-1]{
                continue
            }
            path = append(path,nums[i])
            backtrack(i+1,path)
            path = path[:len(path)-1]
            
        }
    }

    sort.Ints(nums)
    backtrack(0,[]int{})
    return res
}