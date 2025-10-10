/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    maxi := math.MinInt32
	maxPathDown(root, &maxi)
	return maxi
}

func maxPathDown(root *TreeNode, maxi *int) int {
	if root == nil {
		return 0
	}

	lh := max(0, maxPathDown(root.Left, maxi))
	rh := max(0, maxPathDown(root.Right, maxi))

	*maxi = max(*maxi, lh+rh+root.Val)

	return root.Val + max(lh, rh)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}