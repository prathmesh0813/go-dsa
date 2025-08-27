func uniqueOccurrences(arr []int) bool {
    freq := make(map[int]int)

    for _,v := range arr {
        freq[v]++
    }

    seen := make(map[int]bool)

    for _, count := range freq{
        if seen[count]{
            return false
        }

        seen[count] = true
    }
    return true
}