func minEatingSpeed(piles []int, h int) int {
    if len(piles) == 0 {
        return 0
    }
    low,high := 1, findMax(piles)
    for low <= high {
        mid := (low+high)/2
    totalHours := calculateHours(piles, mid)
    if totalHours <= h{
        high = mid - 1
       
    }else{
        low = mid + 1
    }
    }
    return low
}

func calculateHours(v []int, mid int) int{
    total := 0
    for _, pile := range v {
        total += (pile+mid-1) / mid
    }
    return total
}

func findMax(v []int) int{
    if len(v)==0{
        return 0
    }
    max := v[0]
    for _,x := range v[1:]{
        if x > max{
            max = x
        }
    }
    return max
}
