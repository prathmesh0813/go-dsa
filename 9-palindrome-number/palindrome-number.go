func isPalindrome(x int) bool {

str := strconv.Itoa(x)
left, right := 0, len(str)-1

if (x < 0){
    return false
}

for left < right {
    if str[left] != str[right]{
        return false
    }
        left++
        right--
}  
    return true
}