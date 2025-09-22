package main

func printBoundary(root *Node) []int {
	res := []int{}

	if root == nil {
		return nil
	}

	if !isLeaf(root) {
		res = append(res, root.Value)
	}

	addLeftBoundary(root, &res)
	addLeaves(root, &res)
	addRightBoundary(root, &res)

	return res

}

func addLeftBoundary(root *Node, res *[]int) {
	curr := root.Left

	for curr != nil {
		if !isLeaf(curr) {
			*res = append(*res, curr.Value)
		}

		if curr.Left != nil {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
}

func addRightBoundary(root *Node, res *[]int) {
	curr := root.Right
	tmp := []int{}
	for curr != nil {
		if !isLeaf(curr) {
			tmp = append(tmp, curr.Value)
		}

		if curr.Right != nil {
			curr = curr.Right
		} else {
			curr = curr.Left
		}
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		*res = append(*res, tmp[i])
	}
}

func addLeaves(root *Node, res *[]int) {
	if isLeaf(root) {
		*res = append(*res, root.Value)
	}

	if root.Left != nil {
		addLeaves(root.Left, res)
	}

	if root.Right != nil {
		addLeaves(root.Right, res)
	}
}

func isLeaf(root *Node) bool {
	return root.Left == nil && root.Right == nil
}
