package main

func buildParentMap(root *Node, parent map[*Node]*Node, target int) *Node {
	if root == nil {
		return nil
	}
	var targetNode *Node
	queue := []*Node{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Value == target {
			targetNode = node
		}

		if node.Left != nil {
			parent[node.Left] = node
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			parent[node.Right] = node
			queue = append(queue, node.Right)
		}
	}

	return targetNode
}

func burnTree(start *Node, parent map[*Node]*Node) int {
	visited := make(map[*Node]bool)
	queue := []*Node{start}
	visited[start] = true
	time := 0

	for len(queue) > 0 {
		size := len(queue)
		burned := false
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil && !visited[node.Left] {
				visited[node.Left] = true
				queue = append(queue, node.Left)
				burned = true
			}

			if node.Right != nil && !visited[node.Right] {
				visited[node.Right] = true
				queue = append(queue, node.Right)
				burned = true
			}

			if p, ok := parent[node]; ok && !visited[p] {
				visited[p] = true
				queue = append(queue, p)
				burned = true
			}
		}
		if burned {
			time++
		}
	}

	return time
}

func minTimeToBurnTree(root *Node, target int) int {
	parent := make(map[*Node]*Node)
	startNode := buildParentMap(root, parent, target)
	return burnTree(startNode, parent)
}
