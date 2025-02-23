func hasSameDigits(s string) bool {
    for len(s)> 2{
    result := ""
        for i :=0 ;i<len(s)-1;i++{
            result +=string( (s[i] + s[i+1])% 10)
        }  
       s=result  
    }

    if len(s) == 2 && s[0] == s[1]{
        return true
    }


    return false
}