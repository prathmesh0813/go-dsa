package main

import "fmt"

func main() {

	//Problem01 preorder traversal Iterative using stack
	// Example tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5

	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	fmt.Println("Preorder Traversal (Iterative using Stack):")
	PreorderIterative(root)
}
