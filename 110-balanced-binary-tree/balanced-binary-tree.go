/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return checkHeight(root) != -1
}

func checkHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := checkHeight(root.Left)
    if left == -1 {
        return -1
    }

    right := checkHeight(root.Right)
    if right == -1 {
        return -1
    }

    if math.Abs(float64(left-right)) > 1 {
        return -1
    }

    if left > right {
        return left + 1
    }
    return right + 1
}