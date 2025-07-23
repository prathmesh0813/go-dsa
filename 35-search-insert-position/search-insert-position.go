func searchInsert(nums []int, target int) int {
    //Solution no 2
        left, right := 0, len(nums)-1
        n := len(nums)
        ans := n
        for left <= right{
            mid := (left+right)/2
            if nums[mid] >= target{
                ans = mid
                right = mid - 1
            }else{
                left = mid + 1
            }
        }
        return ans
    //Solution no 1
    // for left <= right {
    //     mid := left + (right-left)/2
    //     if nums[mid] == target {
    //         return mid
    //     } else if nums[mid] < target {
    //         left = mid + 1
    //     } else {
    //         right = mid - 1
    //     }
    // }

    // // target not found then left is inserted position
    // return left
}