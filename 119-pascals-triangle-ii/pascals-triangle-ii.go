func getRow(rowIndex int) []int {
    row := []int{1}
    ans :=1 
    for i := 1;i<=rowIndex;i++{
        ans = ans * (rowIndex-i+1)
        ans = ans/i
        row = append(row,ans)
    }
    return row
}