func solveNQueens(n int) [][]string {
    var res [][] string

    board := make([][]byte,n)
    for i := 0;i<n;i++{
        board[i] = make([]byte,n)
        for j := 0;j<n;j++{
            board[i][j] = '.'
        }
    }

    // Helper arrays to track occupied rows and diagonals

    leftRow := make([]int,n)
    upperDiagonal := make([]int,2*n-1)
    lowerDiagonal := make([]int,2*n-1)

    var solve func(col int)
    solve = func(col int){
        if col == n {
            // Found a solution; copy board to result
            tmp := make([]string,n)
            for i := 0;i<n;i++{
                tmp[i] = string(board[i])
            }
            res = append(res,tmp)
            return
        }
         for row := 0; row < n; row++ {
            if leftRow[row] == 0 && lowerDiagonal[row+col] == 0 && upperDiagonal[n-1+col-row] == 0 {
                // Place queen
                board[row][col] = 'Q'
                leftRow[row] = 1
                lowerDiagonal[row+col] = 1
                upperDiagonal[n-1+col-row] = 1

                solve(col + 1)

                // Backtrack
                board[row][col] = '.'
                leftRow[row] = 0
                lowerDiagonal[row+col] = 0
                upperDiagonal[n-1+col-row] = 0
            }
         }
    }
     solve(0)
    return res
}