func intersect(nums1 []int, nums2 []int) []int {
    count := make(map[int]int)
    for _,num := range nums1{
        count[num]++
    }

    var res[]int
    for _,num := range nums2{
        if count[num] > 0{
            res = append(res,num)
            count[num]--
        }    
    }
    return res
}