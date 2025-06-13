func intersection(nums1 []int, nums2 []int) []int {
    n1 := len(nums1)
    n2 := len(nums2)

    sort.Ints(nums1)
    sort.Ints(nums2)

    i,j := 0,0
    ans := []int{}

    for i < n1 && j < n2{
        if nums1[i] < nums2[j]{
            i++
        } else if nums2[j] < nums1[i]{
            j++
        }else{
            if len(ans) == 0 || ans[len(ans)-1] != nums1[i] {
                ans = append(ans, nums1[i])
            }
            i++
            j++
        }

    }
    return ans
}