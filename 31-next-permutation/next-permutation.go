func nextPermutation(nums []int)  {
    n := len(nums)
    ind := -1

//Find the first decreasing element from the right
    for i := n-2 ;i>=0;i--{
        if nums[i] < nums[i+1]{
            ind = i
            break
        }
    }

//If no such index, reverse entire array
    if ind == -1{
        reverse(nums,0,n-1)
        return
    }

//Find element just larger than nums[ind] from the right
    for i := n-1 ;i>ind;i--{
        if nums[i] > nums[ind]{
            nums[i],nums[ind] = nums[ind],nums[i]
            break
        }
    }
//Reverse the subarray to the right of ind
    reverse(nums,ind+1,n-1)

}
    func reverse(nums[]int,left,right int){
        for left < right{
            nums[left],nums[right] = nums[right],nums[left]
            left++ 
            right--
        }
    }