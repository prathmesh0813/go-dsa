func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	cole := 1 // This is to track if first column needs to be zeroed

	// Step 1: Mark zeros using first row and column
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			cole = 0
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Step 2: Use the marks to set zeroes
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Step 3: Set the first row to 0 if needed
	if matrix[0][0] == 0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// Step 4: Set the first column to 0 if needed
	if cole == 0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
