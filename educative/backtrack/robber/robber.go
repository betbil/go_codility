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

func rob(root *TreeNode[int]) int {

	// Your code will replace this placeholder return statement

	robbed, notrobbed := robrecur(root)

	return maxx(robbed, notrobbed)
}

func maxx(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func robrecur(root *TreeNode[int]) (int, int) {

	// Your code will replace this placeholder return statement
	if root == nil {
		return 0, 0
	}

	leftrobbed, leftnotrobbed := robrecur(root.Left)
	rightrobbed, rightnotrobbed := robrecur(root.Right)

	robbed := root.Data + leftnotrobbed + rightnotrobbed
	notrobbed := maxx(leftrobbed, leftnotrobbed) + maxx(rightrobbed, rightnotrobbed)

	return robbed, notrobbed
}
