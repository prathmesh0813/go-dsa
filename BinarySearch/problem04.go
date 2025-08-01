package main

func NthRoot(n, m int) int {
	if n <= 0 || m < 0 {
		return -1 // invalid inputs (you can adjust based on definition)
	}

	if m == 0 || m == 1 {
		return m // Root of 0 is 0; root of 1 is 1
	}
	low, high := 1, m
	for low <= high {
		mid := (low + high) / 2
		result := midN(mid, n, m)
		if result == 0 {
			return mid
		} else if result < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func midN(mid, n, m int) int {
	ans := 1
	for i := 0; i < n; i++ {
		if ans > m/mid {
			return 1
		}
		ans *= mid
	}

	if ans == m {
		return 0
	} else if ans < m {
		return -1
	}
	return 1
}
