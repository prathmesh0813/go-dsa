func permute(nums []int) [][]int {
    var res [][] int
    freq := make([]bool,len(nums))

    var backtrack func (path[] int,freq []bool)
    backtrack = func(path [] int, freq []bool){
        if len(path) == len(nums){
            tmp := make([]int,len(path))
            copy(tmp,path)
            res = append(res,tmp)
            return
        }

        for i :=0;i<len(nums);i++{
            if!freq[i]{
                freq[i] = true
                path = append(path,nums[i])
                backtrack(path,freq)
                path = path[:len(path)-1]
                freq[i] = false
            }
        }


    }

    backtrack([]int{},freq)
    return res
}