func sortColors(nums []int)  {
    n := len(nums)
    low,mid,high := 0,0,n-1

    for mid <= high{
        if nums[mid] == 0{
            nums[low],nums[mid] = nums[mid],nums[low]
            low++
            mid++
        } else if nums[mid] == 1{
            mid++
        } else{
            nums[mid],nums[high] = nums[high],nums[mid]
            high--
        }
    }
}