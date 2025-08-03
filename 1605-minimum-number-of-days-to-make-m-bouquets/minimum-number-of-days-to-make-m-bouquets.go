func minDays(bloomDay []int, m int, k int) int {
    if m*k > len(bloomDay){
        return -1
    }
    mini,maxi := bloomDay[0], bloomDay[0]
    for i:=0;i<len(bloomDay);i++{
        mini =  min(mini, bloomDay[i])
        maxi =max(maxi, bloomDay[i])
    }
    low,high := mini,maxi
    for low <= high{
        mid := (low + high)/2
        if possible(bloomDay,mid,m,k){
            high = mid -1
        }else{
            low = mid + 1
        }
    }
    return low
}

func possible(v []int,day,m,k int) bool{
    cnt,noOfB := 0,0
    for i:=0;i<len(v);i++{
        if v[i] <= day{
            cnt++
        }else{
            noOfB += cnt/k
            cnt = 0
        }
    }
    noOfB += cnt/k
    return noOfB >= m
}

func min(a,b int) int{
    if a < b {
        return a
    }
    return b
}

func max(a,b int) int{
   if a > b {
        return a
    }
    return b
}