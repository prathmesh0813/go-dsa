func convertToTitle(columnNumber int) string {
    result := ""

    for columnNumber > 0{
        columnNumber--
        lastChar := rune('A' + columnNumber%26)
        result = string(lastChar) + result
        columnNumber /=26
    }
    return result
}