func rowAndMaximumOnes(mat [][]int) []int {
    n := len(mat)
    if n == 0{
        return []int{-1,0}
    }

    m := len(mat[0])
    maxCount := -1
    index := -1

    for i:=0;i<n;i++{
        sort.Ints(mat[i])
        firstOneIndex := lowerBound(mat[i],1)
        countOnes := m - firstOneIndex

        if countOnes > maxCount{
            maxCount = countOnes
            index = i
        }
    }
    return []int{index,maxCount}
}

func lowerBound(arr[]int, x int) int{
    low,high := 0 ,len(arr)-1
    ans := len(arr)

    for low <= high{
        mid := (low+high)/2
        if arr[mid] >= x{
            ans = mid
            high = mid - 1
        }else{
            low = mid + 1
        }
    }
    return ans
}