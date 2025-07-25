func searchRange(nums []int, target int) []int {
    n:= len(nums)
    first := firstOccurance(nums,n,target)
    if first == -1{
        return []int{-1, -1}
    }
    last := lastOccurance(nums,n,target)
    return []int{first, last}

}
func firstOccurance(nums[]int,n,target int) int{
    low,high := 0,n-1
    first := -1
    for low <= high{
        mid := (low+high)/2
        if nums[mid] ==target{
            first = mid
            high = mid-1
        }else if nums[mid] < target{
            low = mid + 1
        }else{
            high = mid -1
        }
    } 
        return first
}

func lastOccurance(nums[]int,n,target int) int{
    low,high := 0,n-1
    last := -1
    for low <= high{
        mid := (low+high)/2
        if nums[mid] ==target{
            last = mid
            low = mid+1
        }else if nums[mid] < target{
            low = mid + 1
        }else{
            high = mid -1
        }
    } 
        return last
}