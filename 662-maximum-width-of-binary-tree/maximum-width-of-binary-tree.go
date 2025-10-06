/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Pair struct{
    node *TreeNode
    index uint64
}
func widthOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }

    queue := []Pair{{root,0}}
    maxWidth := 0

    for len(queue) > 0{
        size := len(queue)
        firstIndex := queue[0].index
        var lastIndex uint64


       for i := 0 ; i < size ; i++{
        cur := queue[0]
        queue = queue[1:]

        currIndex := cur.index - firstIndex
        node := cur.node

        if i == size - 1{
            lastIndex = currIndex 
        } 


        if node.Left != nil {
            queue = append(queue,Pair{node.Left, 2*currIndex+1} )
        }

        if node.Right != nil {
            queue = append(queue,Pair{node.Right, 2*currIndex+2} )
        }
    }
       width := int(lastIndex - 0 + 1)
       if width > maxWidth{
        maxWidth = width
       }
    }

    return maxWidth
}