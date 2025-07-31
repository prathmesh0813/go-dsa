func mySqrt(x int) int {
//     if x == 0{
//         return 0
//     }

//  r := x
//     for {
//         next := (r + x/r) / 2
//         if next >= r { 
//             return r
//         }
//         r = next
//     }


//Using binary search
low,high := 1,x
for low <= high{
    mid := (low+high)/2
    if (mid*mid) <= x{
        low = mid+1
    }else{
        high = mid -1
    }
}
    return high
}