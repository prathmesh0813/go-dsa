func isPalindrome(s string) bool {
    var cleanstr strings.Builder

    for _, ch := range s{
        if  unicode.IsLetter(ch) || unicode.IsDigit(ch){
            cleanstr.WriteRune(unicode.ToLower(ch))
        }
    }

    str :=  cleanstr.String()
    left, right := 0, len(str)-1

    for left < right {
        if str[left] != str[right]{
            return false
        }
        left++
        right--
    }
    return true
}