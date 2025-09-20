package main

import "math"

func maxPathSum(root *Node) int {
	maxi := math.MinInt32
	maxPathDown(root, &maxi)
	return maxi
}

func maxPathDown(root *Node, maxi *int) int {
	if root == nil {
		return 0
	}

	lh := max(0, maxPathDown(root.Left, maxi))
	rh := max(0, maxPathDown(root.Right, maxi))

	*maxi = max(*maxi, lh+rh+root.Value)

	return root.Value + max(lh, rh)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
