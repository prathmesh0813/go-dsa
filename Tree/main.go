package main

import "fmt"

func main() {

	// //Problem01 preorder traversal Iterative using stack
	// // Example tree:
	// //        1
	// //       / \
	// //      2   3
	// //     / \
	// //    4   5

	// root := &Node{Value: 1}
	// root.Left = &Node{Value: 2}
	// root.Right = &Node{Value: 3}
	// root.Left.Left = &Node{Value: 4}
	// root.Left.Right = &Node{Value: 5}

	// fmt.Println("Preorder Traversal (Iterative using Stack):")
	// PreorderIterative(root)

	//Problem02 Inorder Traversal Iterative using stack
	// Example tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5

	// root := &Node{Value: 1}
	// root.Left = &Node{Value: 2}
	// root.Right = &Node{Value: 3}
	// root.Left.Left = &Node{Value: 4}
	// root.Left.Right = &Node{Value: 5}

	// fmt.Println("Inorder Traversal (Iterative):")
	// InorderIterative(root)

	//Problem03 Postorder Traversal Iterative using two stack
	// Build the tree
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	// root := &Node{Value: 1}
	// root.Left = &Node{Value: 2}
	// root.Right = &Node{Value: 3}
	// root.Left.Left = &Node{Value: 4}
	// root.Left.Right = &Node{Value: 5}

	// fmt.Println("Postorder Traversal (2 Stacks):")
	// fmt.Println(PostorderTwoStacks(root))

	//Problem04 Postorder Traversal Iterative using one stack

	// Build tree
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	// root := &Node{Value: 1}
	// root.Left = &Node{Value: 2}
	// root.Right = &Node{Value: 3}
	// root.Left.Left = &Node{Value: 4}
	// root.Left.Right = &Node{Value: 5}

	// fmt.Println("Postorder Traversal (1 Stack):")
	// fmt.Println(PostorderOneStack(root))

	// //Problem05 Depth of any given binary tree
	// // Example tree:
	// //        1
	// //       / \
	// //      2   3
	// //     / \
	// //    4   5
	// //	     /
	// //		6
	// root := &Node{Value: 1}
	// root.Left = &Node{Value: 2}
	// root.Right = &Node{Value: 3}
	// root.Left.Left = &Node{Value: 4}
	// root.Left.Right = &Node{Value: 5}
	// root.Left.Right.Left = &Node{Value: 6}

	// fmt.Println("Max Depth:", MaxDepth(root))

	//Problem06 Find binary tree is balance or not
	// Example tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5
	// root := &Node{Value: 1}
	// root.Left = &Node{Value: 2}
	// root.Right = &Node{Value: 3}
	// root.Left.Left = &Node{Value: 4}
	// root.Left.Right = &Node{Value: 5}
	// println(IsBalanced(root))

	// // Unbalanced tree
	// root.Left.Left.Left = &Node{Value: 6}
	// root.Left.Left.Left.Left = &Node{Value: 7}

	// println(IsBalanced(root))

	//Problem07 Diameter of binary tree
	// root := &Node{Value: 1}
	// root.Left = &Node{Value: 2}
	// root.Right = &Node{Value: 3}
	// root.Left.Left = &Node{Value: 4}
	// root.Left.Right = &Node{Value: 5}

	// fmt.Println("Diameter:", diameterOfBinaryTree(root))

	//Problem08 MAximum path sum of given binary tree
	/*
	   Example Tree:
	         -10
	         /  \
	        9   20
	           /  \
	          15   7

	   Max path sum = 15 + 20 + 7 = 42
	*/

	// root := &Node{Value: -10}
	// root.Left = &Node{Value: 9}
	// root.Right = &Node{Value: 20}
	// root.Right.Left = &Node{Value: 15}
	// root.Right.Right = &Node{Value: 7}

	// fmt.Println("Maximum Path Sum:", maxPathSum(root))

	//Problem09 Boundary traversal of a binary tree
	/*
	        1
	      /   \
	     2     3
	    / \   / \
	   4   5 6   7
	      / \
	     8   9
	*/
	// root := &Node{Value: 1}
	// root.Left = &Node{Value: 2}
	// root.Right = &Node{Value: 3}
	// root.Left.Left = &Node{Value: 4}
	// root.Left.Right = &Node{Value: 5}
	// root.Right.Left = &Node{Value: 6}
	// root.Right.Right = &Node{Value: 7}
	// root.Left.Right.Left = &Node{Value: 8}
	// root.Left.Right.Right = &Node{Value: 9}

	// result := printBoundary(root)
	// fmt.Println("Boundary Traversal:", result)

	//Problem10 Top view of the binary tree
	// Example Tree:
	//        1
	//      /   \
	//     2     3
	//    / \   / \
	//   4   5 6   7
	// root := &Node{1,
	// 	&Node{2, &Node{4, nil, nil}, &Node{5, nil, nil}},
	// 	&Node{3, &Node{6, nil, nil}, &Node{7, nil, nil}},
	// }

	// fmt.Println(topView(root))

	//Problem11 Bottom View of the Binary Tree
	// Example:
	//      20
	//     /  \
	//    8   22
	//   / \    \
	//  5  3    25
	//    / \
	//   10 14
	// root := &Node{20, nil, nil}
	// root.Left = &Node{8, nil, nil}
	// root.Right = &Node{22, nil, nil}
	// root.Left.Left = &Node{5, nil, nil}
	// root.Left.Right = &Node{3, nil, nil}
	// root.Right.Right = &Node{25, nil, nil}
	// root.Left.Right.Left = &Node{10, nil, nil}
	// root.Left.Right.Right = &Node{14, nil, nil}

	// fmt.Println(bottomView(root))

	//Problem12 Root to Node path of Binary Tree

	//     1
	//    / \
	//   2   3
	//  / \   \
	// 4   5   6
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}
	root.Right.Right = &Node{Value: 6}

	target := 5
	path := []int{}

	if findPath(root, target, &path) {
		fmt.Println("Path from root to node:", path)
	} else {
		fmt.Println("Target not found")
	}
}
