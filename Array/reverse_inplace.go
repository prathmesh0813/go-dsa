package main

import "fmt"

func ReverseInPlace(list []int, start, end int) {
	for i := start; i < end-i+start; i++ {
		list[i], list[end-i+start] = list[end-i+start], list[i]
	}

	fmt.Println("Start Index: ", start, "End Index: ", end, "List: ", list)
}
