package main

import "container/list"

func bottomView(root *Node) []int {
	if root == nil {
		return []int{}
	}

	hdMap := make(map[int]int)
	q := list.New()
	q.PushBack(Pair{root, 0})

	minHd, maxHd := 0, 0

	for q.Len() > 0 {
		front := q.Front()
		q.Remove(front)
		p := front.Value.(Pair)
		node := p.Node
		hd := p.Hd

		hdMap[hd] = node.Value

		if node.Left != nil {
			q.PushBack(Pair{node.Left, hd - 1})
		}

		if node.Right != nil {
			q.PushBack(Pair{node.Right, hd + 1})
		}

		if hd < minHd {
			minHd = hd
		}

		if hd > maxHd {
			maxHd = hd
		}
	}

	res := make([]int, 0, maxHd-minHd+1)
	for h := minHd; h <= maxHd; h++ {
		res = append(res, hdMap[h])
	}
	return res
}
