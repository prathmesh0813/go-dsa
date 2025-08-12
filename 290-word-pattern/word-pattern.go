func wordPattern(pattern string, s string) bool {
    words := strings.Fields(s)

    if len(pattern) != len(words){
        return false
    }

    pMap := make(map[byte]int)
    wMap := make(map[string]int)

    for i:=0;i<len(pattern);i++{
        pChar := pattern[i]
        word := words[i]

        if pMap[pChar] != wMap[word]{
            return false
        }
        pMap[pChar] = i+1
        wMap[word] = i+1
    }
    return true
}