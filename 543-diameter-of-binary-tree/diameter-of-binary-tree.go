/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0
    height(root,&diameter)
    return diameter
}

func height(root *TreeNode, diameter *int) int{
    if root == nil{
        return 0
    }

    lh := height(root.Left,diameter)
    rh := height(root.Right,diameter)

    if lh + rh > *diameter{
        *diameter = lh + rh
    }

    if lh > rh{
        return 1 + lh
    }

    return 1 + rh
}