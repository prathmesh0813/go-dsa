package main

func matrixMedian(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])
	low, high := matrix[0][0], matrix[0][m-1]
	for i := 0; i < n; i++ {
		if matrix[i][0] < low {
			low = matrix[i][0]
		}
		if matrix[i][m-1] > high {
			high = matrix[i][m-1]
		}
	}
	target := (n*m + 1) / 2

	for low < high {
		mid := (low + high) / 2
		if countSmallEqual(matrix, mid) >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func countSmallEqual(matrix [][]int, x int) int {
	cnt := 0
	for _, row := range matrix {
		cnt += upperBound(row, len(row), x)
	}
	return cnt
}

func upperBound(nums []int, n, target int) int {
	low, high := 0, n-1
	ans := n
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] > target {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}
