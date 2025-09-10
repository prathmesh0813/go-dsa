package main

import "fmt"

func InorderIterative(root *Node) []int {
	curr := root
	out := []int{}
	stack := []*Node{}

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		out = append(out, curr.Value)
		fmt.Print(curr.Value, " ")

		curr = curr.Right
	}
	return out
}
