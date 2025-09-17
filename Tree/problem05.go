package main

func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	if leftDepth > rightDepth {
		return 1 + leftDepth
	}

	return 1 + rightDepth
}
