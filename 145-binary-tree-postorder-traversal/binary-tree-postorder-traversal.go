/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    stack1 := []*TreeNode{root}
    stack2 := []*TreeNode{}
    result := []int{}

    for len(stack1) > 0 {
        curr := stack1[len(stack1)-1]
        stack1 = stack1[:len(stack1)-1]

        stack2 = append(stack2,curr)

        if curr.Left != nil {
            stack1 = append(stack1, curr.Left)
        }

        if curr.Right != nil {
            stack1 = append(stack1, curr.Right)
        }
    }

    for i := len(stack2)-1; i>=0 ;i--{
        result = append(result, stack2[i].Val)
    }
    return result
}