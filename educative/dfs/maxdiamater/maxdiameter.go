package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

// todobetul check anlamadım
func chatgbtdiameterOfBinaryTree(root *TreeNode) int {
	ans = 1
	depth(root)
	return ans - 1
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	L := depth(node.Left)
	R := depth(node.Right)
	ans = max(ans, L+R+1)
	return max(L, R) + 1
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(chatgbtdiameterOfBinaryTree(root)) // Output: 3
}
