func missingNumber(nums []int) int {
    xor1, xor2 := 0,0
    n := len(nums)
    for i := 0; i<n;i++{
        xor1 ^= i
        xor2 ^= nums[i]
    }
    xor1 ^= n
    return xor1 ^ xor2
}