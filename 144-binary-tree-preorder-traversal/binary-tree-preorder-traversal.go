/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	out := []int{}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		out = append(out, curr.Val)
		fmt.Print(curr.Val, " ")

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return out
}