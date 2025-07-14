func reversePairs(nums []int) int {
    return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(arr []int,low,high int) int{
    if low >=high{
        return 0
    }

    mid := (low+high)/2
    count := mergeSort(arr,low,mid)
    count += mergeSort(arr,mid + 1,high)
    count += merge(arr,low,mid,high)
    return count
}

func merge(arr []int,low,mid,high int) int{
    count := 0
    j := mid+1

    for i := low;i<=mid;i++{
        for j <= high && arr[i]>2*arr[j]{
            j++
        }
        count += j - (mid+1)
    }

    temp := make([]int,0,high-low+1)
    left := low
    right := mid+1

    for left<=mid && right<= high{
        if arr[left] <= arr[right]{
            temp = append(temp,arr[left])
            left++
        }else{
            temp = append(temp,arr[right])
            right++
        }
    }

    for left <= mid{
        temp = append(temp,arr[left])
        left++
    }
    for right <= high{
        temp = append(temp,arr[right])
        right++
    }
    for i := low; i <= high; i++ {
        arr[i] = temp[i-low]
    }
    return count

}