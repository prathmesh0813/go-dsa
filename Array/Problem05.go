package main

func alternateNumbers(a []int) []int {
	pos := []int{}
	neg := []int{}

	for _, val := range a {
		if val > 0 {
			pos = append(pos, val)
		} else {
			neg = append(neg, val)
		}
	}

	//n := len(a)
	if len(pos) > len(neg) {
		for i := 0; i < len(neg); i++ {
			a[2*i] = pos[i]
			a[2*i+1] = neg[i]
		}
		index := len(neg) * 2
		for i := len(neg); i < len(pos); i++ {
			a[index] = pos[i]
			index++
		}
	} else {
		for i := 0; i < len(pos); i++ {
			a[2*i] = pos[i]
			a[2*i+1] = neg[i]
		}
		index := len(pos) * 2
		for i := len(pos); i < len(neg); i++ {
			a[index] = neg[i]
			index++
		}
	}
	return a
}
