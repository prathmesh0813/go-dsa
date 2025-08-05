func shipWithinDays(weights []int, days int) int {
    low := weights[0]
    sum := 0
    for _,weight := range weights{
        if weight > low{
            low = weight
        }
        sum += weight
    }
    high := sum

    for low <= high{
        mid := (low+high)/2
        if findDays(weights,mid) <= days{
            high = mid - 1
        }else{
            low = mid + 1
        }
    }
    return low
}

func findDays(weights []int,capacity int) int {
    days := 1
    load := 0
    for _,weight := range weights{
        if load+weight > capacity{
            days++
            load = weight
        }else{
            load += weight
        }
    }
    return days
}