func spiralOrder(matrix [][]int) []int {
    n := len(matrix)
    m := len(matrix[0])

    left,right := 0,m-1
    top,bottom := 0,n-1
    ans := []int{}

    for left <= right && top <= bottom{
        for i := left;i <= right;i++{
            ans = append(ans,matrix[top][i])
        }
        top++
        for i := top ; i <= bottom;i++{
            ans = append(ans,matrix[i][right])
        }
        right--
        if top <= bottom{
            for i := right;i >= left;i--{
                ans = append(ans,matrix[bottom][i])
            }
        bottom--
        }
    if left <= right{
        for i := bottom;i >= top ;i--{
            ans = append(ans,matrix[i][left])
        }
    left++
    }
    }
    return ans
}