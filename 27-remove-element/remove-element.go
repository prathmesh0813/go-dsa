func removeElement(nums []int, val int) int {
    count :=0
   for _,item :=range nums{
        if item != val{
            nums[count]=item
            count ++
        }
    }
    return count
}