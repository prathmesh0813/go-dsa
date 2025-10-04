/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    result := []int{}
    rightView(root,&result,0)
    return result
}

func rightView(root *TreeNode,result *[]int,currDepth int){
    if root == nil{
        return
    }

    if currDepth == len(*result){
        *result = append(*result,root.Val)
    }

    rightView(root.Right,result,currDepth+1)
    rightView(root.Left,result,currDepth+1)
}