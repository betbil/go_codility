// Definition of a binary tree node
// type TreeNode[T any] struct {
// 	Data  T
// 	Left  *TreeNode[T]
// 	Right *TreeNode[T]
// }

package main

type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func mirrorBinaryTree(root *TreeNode[int]) *TreeNode[int] {
	// Replace this placeholder return statement with your code
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorBinaryTree(root.Left)
	mirrorBinaryTree(root.Right)

	return root
}
