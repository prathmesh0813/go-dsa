package main

import "math"

func BalancedBinaryTree(root *Node) int {
	if root == nil {
		return 0
	}

	leftDepth := BalancedBinaryTree(root.Left)
	rightDepth := BalancedBinaryTree(root.Right)

	if leftDepth == -1 {
		return -1
	}
	if rightDepth == -1 {
		return -1
	}
	if math.Abs(float64(leftDepth)-float64(rightDepth)) > 1 {
		return -1
	}

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}

func IsBalanced(root *Node) bool {
	return BalancedBinaryTree(root) != -1
}
