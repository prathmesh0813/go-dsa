func combinationSum(candidates []int, target int) [][]int {
    var res [][] int

    var backtrack func(ind int, path []int, total int)
    backtrack = func(ind int,path []int,total int){
        if ind == len(candidates){
            if total == target{
                comb := make([]int,len(path))
                copy(comb,path)
                res = append(res,comb)
            }
            return
        }
        if total+candidates[ind] <= target{
            path = append(path,candidates[ind])
            backtrack(ind, path, total+candidates[ind])
            path = path[:len(path)-1]
        }
        backtrack(ind+1, path,total)
    }



    backtrack(0,[]int{},0)
    return res
}