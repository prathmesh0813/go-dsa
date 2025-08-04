func smallestDivisor(nums []int, threshold int) int {
    if len(nums) > threshold{
        return -1
    }
    low,high := 1,max(nums)
    for low <= high{
        mid := (low+high)/2
        if sumByD(nums,mid) <= threshold{
            high = mid - 1
        }else{
            low = mid +1 
        }
    }
    return low
}

func sumByD(nums[]int, mid int) int {
    sum := 0
    for _,v := range nums{
        sum += int(math.Ceil(float64(v)/float64(mid)))
    }
    return sum
}

func max(nums[]int) int{
    max := nums[0]
    for _,v := range nums[1:]{
        if v > max{
            max = v
        }
    }
    return max
}