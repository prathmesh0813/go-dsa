func partition(s string) [][]string {
    var res [][] string 

    var backtrack func(index int, path []string)
    backtrack = func(index int, path []string){
    if index == len(s){
        tmp := make([]string,len(path))
        copy(tmp,path)
        res = append(res,tmp)
        return
    }
    for i := index;i<len(s);i++{
        if isPalindrome(s,index,i){
            path = append(path,s[index:i+1])
            backtrack(i+1,path)
            path = path[:len(path)-1]
        }
    }
    }

    backtrack(0,[]string{})
    return res
}

func isPalindrome(s string, start int, end int) bool{
    for start < end{
        if s[start] != s[end]{
            return false
        }

        start++
        end--
    }
    return true
}