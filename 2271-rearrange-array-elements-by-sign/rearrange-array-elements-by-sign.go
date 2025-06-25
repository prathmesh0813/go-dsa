func rearrangeArray(nums []int) []int {
    positive,negative := 0,1
    ans := make([]int,len(nums))
    for i := 0;i < len(nums);i++{
        if nums[i] < 0{
            ans[negative] = nums[i]
            negative += 2
        }else {
            ans[positive] = nums[i]
            positive += 2
        }
    }
    return ans
}