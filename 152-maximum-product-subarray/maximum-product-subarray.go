func maxProduct(nums []int) int {
    prefix,suffix := 1,1
    ans := nums[0]
    n := len(nums)

    for i := 0 ; i < n ;i++{
        if prefix == 0{
            prefix = 1
        }
        if suffix == 0{
            suffix = 1
        }
        prefix = prefix * nums[i]
        suffix = suffix * nums[n-i-1]
        ans = max(ans,prefix,suffix)
    }
    return ans
}

func max(a, b, c int) int {
    if a > b {
        if a > c {
            return a
        }
        return c
    } else {
        if b > c {
            return b
        }
        return c
    }
}