/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }

    lh := getHeightLeft(root)
    rh := getHeightRight(root)

    if lh == rh{
        return (1 << lh) - 1
    }else{
        return countNodes(root.Left) + countNodes(root.Right) + 1
    }
}

func getHeightLeft(root *TreeNode) int {
    count := 0
    for root != nil {
        count++
        root = root.Left
    }
    return count
}

func getHeightRight(root *TreeNode) int {
    count := 0
    for root != nil {
        count++
        root = root.Right
    }
    return count
}