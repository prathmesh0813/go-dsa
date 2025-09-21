/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    result := [][]int{}
    if root == nil{
        return result
    }
    queue := []*TreeNode{root}
    leftToRight := true

    for len(queue) > 0{
        size := len(queue)
        row := make([]int,size)

        for i:=0 ; i<size; i++{
            node := queue[0]
            queue = queue[1:]

            var index int
            if leftToRight {
                index = i
            }else{
                index = size - 1 - i
            }

            row[index] = node.Val

            if node.Left != nil {
                queue = append(queue,node.Left)
            }

            if node.Right != nil {
                queue = append(queue,node.Right)
            }
        }   

        result = append(result,row)
        leftToRight = !leftToRight
    }
    return result
}