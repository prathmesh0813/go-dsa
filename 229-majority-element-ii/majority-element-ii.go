func majorityElement(nums []int) []int {
   var count1,count2,el1,el2 int
    for i := 0 ;i<len(nums);i++{
        if count1 == 0 && el2 != nums[i]{
            count1 = 1
            el1 = nums[i]
        }else if count2 == 0 && el1 != nums[i]{
            count2 = 1
            el2 = nums[i]
        }else if nums[i] == el1{
            count1++
        }else if nums[i] == el2{
            count2++
        }else{
            count1--
            count2--
        }
    }
    count1,count2 = 0,0
    for _,num := range nums{
        if num == el1{
            count1++
        }else if num == el2{
            count2++
        }
    }
    result := []int{}
    n := len(nums)
    if count1 > n/3{
        result = append(result,el1)
    }
    if count2 > n/3 {
        result = append(result, el2)
    }
    return result

    }