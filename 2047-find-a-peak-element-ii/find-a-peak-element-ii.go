func findPeakGrid(mat [][]int) []int {
    n := len(mat)
    m := len(mat[0])

    low,high := 0,m-1
    for low <= high{
        mid := (low+high)/2
        maxRowIndex := findMaxIndex(mat,n,m,mid)
        left := -1
        if mid - 1 >= 0{
            left = mat[maxRowIndex][mid-1]
        }
        right := -1
        if mid + 1 < m{
            right = mat[maxRowIndex][mid+1]
        }

        if mat[maxRowIndex][mid] > left && mat[maxRowIndex][mid] > right{
            return []int{maxRowIndex,mid}
        }else if mat[maxRowIndex][mid] < left {
            high = mid-1
        }else{
            low = mid + 1
        }
    }
    return []int{-1,-1}
}

func findMaxIndex(mat [][]int, n, m, col int) int {
	maxValue := -1
	index := -1
	for i := 0; i < n; i++ {
		if mat[i][col] > maxValue {
			maxValue = mat[i][col]
			index = i
		}
	}
	return index
}