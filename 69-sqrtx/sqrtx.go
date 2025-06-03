func mySqrt(x int) int {
    if x == 0{
        return 0
    }

 r := x
    for {
        next := (r + x/r) / 2
        if next >= r { 
            return r
        }
        r = next
    }
}