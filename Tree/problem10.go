package main

import (
	"container/list"
	"sort"
)

type Pair struct {
	Node *Node
	Hd   int
}

func topView(root *Node) []int {
	result := []int{}
	if root == nil {
		return result
	}
	hdMap := make(map[int]int)
	queue := list.New()
	queue.PushBack(Pair{root, 0})

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		item := front.Value.(Pair)
		node, hd := item.Node, item.Hd

		if _, exits := hdMap[hd]; !exits {
			hdMap[hd] = node.Value
		}

		if node.Left != nil {
			queue.PushBack(Pair{node.Left, hd - 1})
		}

		if node.Right != nil {
			queue.PushBack(Pair{node.Right, hd + 1})
		}
	}
	keys := make([]int, 0, len(hdMap))
	for k := range hdMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		result = append(result, hdMap[k])
	}
	return result
}
