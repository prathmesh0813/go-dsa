func findMin(nums []int) int {
    n := len(nums)
    low,high := 0,n-1
    ans := math.MaxInt

    for low <= high{
        mid := (low+high)/2
        if nums[low] <= nums[high]{
            ans = min(ans,nums[low])
            break
        }
        if nums[low] <= nums[mid]{
            ans = min(ans,nums[low])
            low = mid+1
        }else{
            ans=min(ans,nums[mid])
            high=mid-1
        }
    }
    return ans
}

func min(a,b int) int{
    if a > b{
        return b
    }
    return a
}