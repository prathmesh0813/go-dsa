func splitArray(nums []int, k int) int {
    if k > len(nums){
        return -1
    }
    low := max(nums)
    high := sum(nums)

    for low <= high {
        mid := (low + high)/2
       students := countStudents(nums,mid)
       if students > k{
        low = mid + 1
       }else {
        high = mid - 1
       }
    }
    return low
}

func countStudents(nums []int, pages int) int {
    students := 1
    pagesStudent := 0
    for _, val := range nums{
        if pagesStudent+val <= pages{
            pagesStudent += val
        }else{
            students++
            pagesStudent = val
        }
    }
    return students
}

func sum(nums[] int) int {
    total := 0
    for _, v := range nums{
        total += v
    }
    return total
}

func max(nums[] int) int {
    maxVal := nums[0]
    for _,v := range nums{
        if v > maxVal{
            maxVal = v
        }
    }
    return maxVal
}