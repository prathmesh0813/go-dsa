func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    totalNums := n * n
    var sum, sumSq int64
    expectedSum := int64(totalNums*(totalNums+1))/2
    expectedSumSq := int64(totalNums*(totalNums+1)*(2*totalNums+1)) / 6

    for i := 0 ; i < n ;i++{
        for j := 0;j < n ;j++{
            val := int64(grid[i][j])
            sum += val
            sumSq += val * val
        }
    }

    diff := sum - expectedSum
    sqdiff := sumSq - expectedSumSq

    sumXY := sqdiff/diff

    x:=(diff + sumXY)/2
    y := x-diff

    return []int{int(x),int(y)}
}