func longestConsecutive(nums []int) int {
   
    if len(nums)==0{
        return 0
    }

    st := make(map[int]bool)
    for _,num := range nums{
        st[num] = true
    }

    longest := 1

    for num := range st{
        if !st[num-1]{
            cnt := 1 
            current := num    
    
        for st[current+1]{
            cnt++
            current++
        }

        if cnt > longest{
            longest = cnt
        }
    }
}
    return longest
}