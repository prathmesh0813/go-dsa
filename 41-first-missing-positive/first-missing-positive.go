func firstMissingPositive(nums []int) int {

    n :=len(nums)

    for i:=0;i<n;{
        correctIndex := nums[i]-1

        if nums[i] >0 && nums[i]<=n && nums[i] != nums[correctIndex]{
            nums[i],  nums[correctIndex] = nums[correctIndex],nums[i]
        }else{
            i++
        }
    }

        for i:=0;i<n;i++{
            if nums[i] != i+1{
                return i+1
            }
        }
    return n+1
}