func findTheDifference(s string, t string) byte {
  sum1:=0
  sum2:=0

    for i := 0; i < len(s); i++ {
        sum1 += int(s[i])
    }

    for i := 0; i < len(t); i++ {
        sum2 += int(t[i])
    }

    return byte(sum2 - sum1)
}
