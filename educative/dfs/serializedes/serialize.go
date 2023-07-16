// Definition of a binary tree node
// type TreeNode[T any] struct {
// 	Data  T
// 	Left  *TreeNode[T]
// 	Right *TreeNode[T]
// }

package main

import (
	"fmt"
	"strconv"
)

type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func serialize(root *TreeNode[int]) []string {
	// Replace this placeholder return statement with your code
	serializeList := make([]string, 0)
	if root == nil {
		return serializeList
	}
	preorderTraversalSerialize(root, &serializeList)
	return serializeList
}

func preorderTraversalSerialize(root *TreeNode[int], serializelList *[]string) {
	// Your code here
	if root == nil {
		*serializelList = append(*serializelList, "*")
		return
	}
	*serializelList = append(*serializelList, fmt.Sprintf("%d", root.Data))
	preorderTraversalSerialize(root.Left, serializelList)
	preorderTraversalSerialize(root.Right, serializelList)

	return
}

func deserialize(stream *[]string) *TreeNode[int] {

	// Replace this placeholder return statement with your code
	if len(*stream) == 0 {
		return nil
	}

	stack := make([]*TreeNode[int], 0)
	root := &TreeNode[int]{}
	crnode := root
	var tmpNode *TreeNode[int]
	for _, v := range *stream {
		tmpNode = nil
		if v != "*" {
			tmpNode = &TreeNode[int]{}
			i, _ := strconv.Atoi(v)
			tmpNode.Data = i
		}
		if crnode == nil {
			crnode = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			crnode.Right = tmpNode
			crnode = crnode.Right
		} else {
			crnode.Left = tmpNode
			crnode = crnode.Left
		}

		if crnode != nil {
			stack = append(stack, crnode)
		}
	}

	return root.Left
}

// todobetul check güzel cözüm
func edudeserialize(stream *[]string) *TreeNode[int] {
	// dequeue first element form array
	val := (*stream)[0]
	fmt.Println("val:", val)
	*stream = (*stream)[1:len(*stream)] // reframing the slice takes O(1) time complexity

	// Return nil when a marker is encountered
	if string(val[0]) == "*" {
		return nil
	}

	// Creating new Binary Tree Node from current value from stream
	tempVal, _ := strconv.Atoi(val)
	node := &TreeNode[int]{Data: tempVal}

	// Doing a pre-order tree traversal for serialization
	node.Left, node.Right = edudeserialize(stream), edudeserialize(stream)

	// Return node if it exists
	return node
}

func deneme(serializelList *[]string) {
	// Your code here
	*serializelList = append(*serializelList, "oley")
	return
}

func main() {
	fmt.Println("deneme")
	serializeList2 := make([]string, 0)
	fmt.Println("before deneme serializeList2:", serializeList2)
	deneme(&serializeList2)
	fmt.Println("after deneme serializeList2:", serializeList2)

	node5 := &TreeNode[int]{Data: 5, Left: nil, Right: nil}
	node4 := &TreeNode[int]{Data: 4, Left: nil, Right: nil}
	node19 := &TreeNode[int]{Data: 19, Left: nil, Right: nil}
	node1 := &TreeNode[int]{Data: 1, Left: nil, Right: nil}
	node17 := &TreeNode[int]{Data: 17, Left: node19, Right: node5}
	node2 := &TreeNode[int]{Data: 2, Left: node1, Right: node4}
	node3 := &TreeNode[int]{Data: 3, Left: node2, Right: node17}

	// Print the tree
	serial := serialize(node3)
	fmt.Println("serial:", serial)
	deserial := deserialize(&serial)
	fmt.Println("deserial:", deserial)

	edudeserialize := edudeserialize(&serial)
	fmt.Println("edudeserialize:", edudeserialize)
}
