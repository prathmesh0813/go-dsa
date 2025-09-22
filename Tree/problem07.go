package main

func height(root *Node, diameter *int) int {
	if root == nil {
		return 0
	}

	lh := height(root.Left, diameter)
	rh := height(root.Right, diameter)

	if lh+rh > *diameter {
		*diameter = lh + rh
	}

	if lh > rh {
		return 1 + lh
	}
	return 1 + rh
}

func diameterOfBinaryTree(root *Node) int {
	diameter := 0
	height(root, &diameter)
	return diameter
}
