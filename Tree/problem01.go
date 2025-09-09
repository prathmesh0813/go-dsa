package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func PreorderIterative(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	out := []int{}

	for len(stack) > 0 {
		// pop top
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// visit
		out = append(out, curr.Value)
		fmt.Print(curr.Value, " ")

		// push right then left so left is processed first
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return out
}
