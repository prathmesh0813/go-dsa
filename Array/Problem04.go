package main

func UnionOfArray(a, b []int) []int {
	n1 := len(a)
	n2 := len(b)

	i, j := 0, 0
	unionarr := []int{}

	for i < n1 && j < n2 {
		if a[i] < b[j] {
			if len(unionarr) == 0 || unionarr[len(unionarr)-1] != a[i] {
				unionarr = append(unionarr, a[i])
			}
			i++
		} else {
			if len(unionarr) == 0 || unionarr[len(unionarr)-1] != b[j] {
				unionarr = append(unionarr, b[j])
			}
			j++
		}
	}

	for i < n1 {
		if len(unionarr) == 0 || unionarr[len(unionarr)-1] != a[i] {
			unionarr = append(unionarr, a[i])
		}
		i++
	}

	for j < n2 {
		if len(unionarr) == 0 || unionarr[len(unionarr)-1] != b[j] {
			unionarr = append(unionarr, b[j])
		}
		j++
	}
	return unionarr
}
