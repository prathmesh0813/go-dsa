func searchMatrix(matrix [][]int, target int) bool {
    n := len(matrix)
    m := len(matrix[0])
    row,col:=0,m-1
    for row < n && col >= 0{
        if matrix[row][col] == target{
            return true
        }else if matrix[row][col] < target{
            row++
        }else{
            col--
        }
    }
    return false
}