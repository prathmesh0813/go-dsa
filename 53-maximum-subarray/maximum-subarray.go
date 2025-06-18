func maxSubArray(nums []int) int {
    sum, maxi := 0,math.MinInt64
    for i := 0 ; i < len(nums) ; i++{
        sum += nums[i] 
        if sum > maxi{
            maxi = sum
        }

        if sum < 0{
            sum = 0
        }
    }
    return maxi
}