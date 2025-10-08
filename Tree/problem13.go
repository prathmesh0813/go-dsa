package main

import "fmt"

func changeTree(root *Node) {
	if root == nil {
		return
	}

	childSum := 0

	if root.Left != nil {
		childSum += root.Left.Value
	}
	if root.Right != nil {
		childSum += root.Right.Value
	}

	if childSum >= root.Value {
		root.Value = childSum
	} else {
		if root.Left != nil {
			root.Left.Value = root.Value
		}
		if root.Right != nil {
			root.Right.Value = root.Value
		}
	}

	changeTree(root.Left)
	changeTree(root.Right)

	total := 0
	if root.Left != nil {
		total += root.Left.Value
	}
	if root.Right != nil {
		total += root.Right.Value
	}

	if root.Left != nil || root.Right != nil {
		root.Value = total
	}
}

//to print preorder traversal
func preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Value, " ")
	preorder(root.Left)
	preorder(root.Right)
}
