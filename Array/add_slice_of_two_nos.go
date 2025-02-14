package main

import "fmt"

func AddSliceOfTwoNos(num1, num2 []int) {

	num1, num2 = equlizeSliceLength(num1, num2)
	carry := false

	for i := len(num1) - 1; i >= 0; i-- {
		num1[i] += num2[i]

		if carry {
			num1[i]++
			carry = false
		}
		if num1[i] >= 10 {
			num1[i] -= 10
			carry = true
		}
	}

	if carry {
		num1 = append([]int{1}, num1...)
	}
	fmt.Println(num1)

}

func equlizeSliceLength(num1, num2 []int) ([]int, []int) {

	diff := diffOfSliceLength(len(num1), len(num2))

	zeros := make([]int, diff)

	if len(num2) > len(num1) {
		num1 = append(num1, zeros...)
	} else if len(num1) > len(num2) {
		num2 = append(num2, zeros...)
	}
	return num1, num2

}

func diffOfSliceLength(a, b int) int {

	if b > a {
		return b - a
	}
	return a - b

}
