func reverse(x int) int {
    rev := 0
    for x!=0{
        digit := x % 10
        x /= 10
        if rev > math.MaxInt32/10 || rev < math.MinInt32/10{
            return 0
        }
        rev = rev*10 + digit
    }
    return rev
}