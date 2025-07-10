func fourSum(nums []int, target int) [][]int {
     sort.Ints(nums)
     var result [][] int
     n := len(nums)

     for i := 0;i<n;i++{
        if i>0 && nums[i] == nums[i-1]{
            continue
        }
        for j := i+1;j<n;j++{
            if j != i+1 && nums[j]==nums[j-1]{
                continue
            }
            k := j+1
            l := n-1
            for k < l{
                sum := nums[i]
                sum += nums[j]
                sum += nums[k]
                sum += nums[l]
                if sum == target{
                    result = append(result,[]int {nums[i],nums[j],nums[k],nums[l]})
                    k++
                    l--
                    for k < l && nums[k]==nums[k-1]{
                        k++
                    }
                    for k < l && nums[l] == nums[l+1]{
                        l--
                    }
                }else if sum < target{
                    k++
                }else{
                    l--
                }
            }
        }
     }
     return result
}