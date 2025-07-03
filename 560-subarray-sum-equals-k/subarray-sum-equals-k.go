func subarraySum(nums []int, k int) int {
    result := make(map[int]int)
    result[0]=1

    preSum,cnt := 0,0

    for i := 0; i<len(nums);i++{
        preSum += nums[i]
        remove := preSum - k
        cnt += result[remove]
        result[preSum] += 1
    }
    return cnt
}