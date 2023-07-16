// Definition of a binary tree node
// type TreeNode[T any] struct {
// 	Data  T
// 	Left  *TreeNode[T]
// 	Right *TreeNode[T]
// }

package main

import (
	"container/list"
	"fmt"
)

type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func flattenBinaryTree(root *TreeNode[int]) *TreeNode[int] {
	// your code will replace this placeholder return statement
	if root == nil {
		return nil
	}

	nodeStk := list.List{}
	nodeStk.PushBack(root)
	retNode := &TreeNode[int]{Data: root.Data}
	returnNode := retNode
	crnode := root.Left

	for nodeStk.Len() > 0 || crnode != nil {
		for crnode != nil {
			retNode.Right = &TreeNode[int]{Data: crnode.Data}
			retNode = retNode.Right
			nodeStk.PushBack(crnode)
			crnode = crnode.Left
		}

		if nodeStk.Len() > 0 {
			popedNode := nodeStk.Back()
			crnode = popedNode.Value.(*TreeNode[int]).Right
			nodeStk.Remove(popedNode)
		}
	}

	return returnNode
}

// todobetul check güzel cözüm
func eduflattenBinaryTree(root *TreeNode[int]) *TreeNode[int] {
	if root == nil {
		return nil
	}
	current := root
	for current != nil {

		if current.Left != nil {
			last := current.Left
			for last.Right != nil {
				last = last.Right
			}

			last.Right = current.Right
			current.Right = current.Left
			current.Left = nil
		}

		current = current.Right
	}
	return root
}

func printTree[T any](node *TreeNode[T]) {
	if node == nil {
		return
	}

	fmt.Print(node.Data, ",")
	printTree(node.Left)
	printTree(node.Right)
}

func main() {
	// your code goes here
	node5 := &TreeNode[int]{Data: 5, Left: nil, Right: nil}
	node4 := &TreeNode[int]{Data: 4, Left: nil, Right: nil}
	node19 := &TreeNode[int]{Data: 19, Left: nil, Right: nil}
	node1 := &TreeNode[int]{Data: 1, Left: nil, Right: nil}
	node17 := &TreeNode[int]{Data: 17, Left: node19, Right: node5}
	node2 := &TreeNode[int]{Data: 2, Left: node1, Right: node4}
	node3 := &TreeNode[int]{Data: 3, Left: node2, Right: node17}

	// Print the tree
	printTree(node3)

	fmt.Println("Flatten Binary Tree")
	noden := flattenBinaryTree(node3)
	printTree(noden)
	//[3,2,17,1,4,19,5]

}
