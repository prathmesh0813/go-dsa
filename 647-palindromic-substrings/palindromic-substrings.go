func countSubstrings(s string) int {

    count :=0 
    expandCount := func (left, right int){
        for left >=0 && right < len(s) && s[left] == s[right]{
            count++
            left--
            right++
        }
    }
    for i:=0;i<len(s);i++{
        expandCount(i,i)

        expandCount(i,i+1)
    }
    return count
}