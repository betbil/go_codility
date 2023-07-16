// Definition of a binary tree node
// type TreeNode[T any] struct {
// 	Data  T
// 	Left  *TreeNode[T]
// 	Right *TreeNode[T]
// }

package main

import (
	"strconv"
)

type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func levelOrderTraversal(root *TreeNode[int]) string {
	// Write your code here
	// your code will replace this placeholder return statement
	if root == nil {
		return ""
	}
	queue := make([]*TreeNode[int], 0)
	queue = append(queue, root)
	depth := 0
	outstr := make([]byte, 0)
	for len(queue) > 0 {
		if depth > 0 {
			outstr = append(outstr, []byte(" : ")...)
		}
		depth++
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i > 0 {
				outstr = append(outstr, []byte(", ")...)
			}
			outstr = append(outstr, []byte(strconv.Itoa(node.Data))...)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
	}
	return string(outstr)
}
