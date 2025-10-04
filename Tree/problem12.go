package main

func findPath(root *Node, target int, path *[]int) bool {
	if root == nil {
		return false
	}

	*path = append(*path, root.Value)

	if root.Value == target {
		return true
	}

	if findPath(root.Left, target, path) || findPath(root.Right, target, path) {
		return true
	}

	*path = (*path)[:len(*path)-1]
	return false
}
