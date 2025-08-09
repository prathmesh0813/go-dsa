func maxDistance(position []int, m int) int {
    sort.Ints(position)
    n := len(position)
    low,high :=1 ,position[n-1] - position[0]

    for low <= high{
        mid := (low + high)/2
        if canWePlace(position,mid,m) == true {
            low = mid + 1
        }else{
            high = mid - 1
        }
    }
    return high
}

func canWePlace(arr [] int , mid , cows int) bool {
    cntCows := 1 
    last := arr[0]
    for i:=1;i<len(arr);i++{
        if arr[i] - last >= mid {
            cntCows++
            last = arr[i]
        }
        if cntCows >= cows {
            return true
        }
    }
    return false
}