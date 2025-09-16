package main

func PostorderOneStack(root *Node) []int {
	result := []int{}
	stack := []*Node{}
	var lastVisited *Node
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		peek := stack[len(stack)-1]

		if peek.Right != nil && lastVisited != peek.Right {
			curr = peek.Right
		} else {
			result = append(result, peek.Value)
			lastVisited = peek
			stack = stack[:len(stack)-1]
		}
	}
	return result
}
