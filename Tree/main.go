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
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	fmt.Println("Diameter:", diameterOfBinaryTree(root))
}
