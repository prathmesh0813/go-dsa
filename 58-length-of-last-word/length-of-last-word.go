func lengthOfLastWord(s string) int {
    count :=0
    n := len(s)

    for i:= n-1;i>=0;i--{
        if s[i] != ' '{
            count++
        }else if count != 0{
            break
        }
    }

    return count
}