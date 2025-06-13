func moveZeroes(nums []int)  {
    n := len(nums)

    j := -1
    for i := 0;i<n;i++{
        if nums[i] == 0{
            j = i
            break
        }
    }
if j == -1{
    return
}

    for i := j+1;i<n;i++{
        if nums[i] != 0{
            nums[i],nums[j] = nums[j],nums[i]
            j++
        }
    }
}