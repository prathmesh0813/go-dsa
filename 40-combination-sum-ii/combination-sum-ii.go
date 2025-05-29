func combinationSum2(candidates []int, target int) [][]int {

    var res [][] int
    sort.Ints(candidates)

    var backtrack func(ind int,path []int,target int)
    backtrack = func(ind int, path []int,target int){
        if target == 0{
            comb := make([]int,len(path))
            copy(comb,path)
            res = append(res,comb)
            return
        }
    
    for i := ind; i< len(candidates);i++ {
        if i > ind && candidates[i] == candidates[i-1]{
            continue
        }

        if candidates[i] > target{
            break
        }
        path = append(path,candidates[i])
        backtrack(i+1,path,target - candidates[i])
        path = path[:len(path)-1]
            }

    }

    backtrack(0,[]int{},target)
    return res
    
}