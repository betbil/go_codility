// Definition of a binary tree node
// type TreeNode[T any] struct {
// 	Data  T
// 	Left  *TreeNode[T]
// 	Right *TreeNode[T]
// }

package main

import "fmt"

type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func isSymmetric(root *TreeNode[int]) bool {

	// Replace this placeholder return statement with your code
	if root == nil {
		return true
	}
	queue := make([]*TreeNode[int], 0)
	crnode := root
	queue = append(queue, crnode)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			crnode = queue[0]
			queue = queue[1:]
			if crnode != nil {
				queue = append(queue, crnode.Left)
				queue = append(queue, crnode.Right)
			}
		}
		if !isSymmetricHelper(queue) {
			return false
		}
	}

	return true
}

func isSymmetricHelper(queue []*TreeNode[int]) bool {
	l := len(queue)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		if queue[i] == nil && queue[j] == nil {
			continue
		}
		if queue[i] == nil || queue[j] == nil {
			return false
		}
		if queue[i].Data != queue[j].Data {
			return false
		}
	}
	return true
}

func main() {
	//[1,2,2,3,4,4,3]
	root := &TreeNode[int]{1, &TreeNode[int]{2, &TreeNode[int]{3, nil, nil}, &TreeNode[int]{4, nil, nil}}, &TreeNode[int]{2, &TreeNode[int]{4, nil, nil}, &TreeNode[int]{3, nil, nil}}}
	fmt.Println(isSymmetric(root))

}
