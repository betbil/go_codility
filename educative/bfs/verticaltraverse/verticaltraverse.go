package main

import (
	"fmt"
	"strings"
)

// todobetul güzel cözüm
// verticalOrder is the solution function
func verticalOrder(root *TreeNode[int]) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	nodeList := make(map[int][]int)
	minColumn := 0
	maxIndex := 0

	// push root into the queue
	queue := new(Queue)
	queue.Enqueue(root, 0)

	// traverse over the nodes in the queue
	for !queue.IsEmpty() {
		node, column := queue.Dequeue()

		if node != nil {
			temp := nodeList[column]
			temp = append(temp, node.Data)
			nodeList[column] = temp

			// get min and max column numbers for the tree
			minColumn = min(minColumn, column)
			maxIndex = max(maxIndex, column)

			// add current node's left and right child in the queue
			queue.Enqueue(node.Left, column-1)
			queue.Enqueue(node.Right, column+1)
		}
	}
	result := make([][]int, 0)
	for x := minColumn; x < maxIndex+1; x++ {
		result = append(result, nodeList[x])
	}
	return result
}

// max function returns the maximum of the integers provided
func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

// min function returns the minimum of the integers provided
func min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

// Driver code
func main() {

	// Create a list of list of TreeNode objects to represent binary trees
	listOfTrees := [][]*TreeNode[int]{
		{&TreeNode[int]{Data: 100}, &TreeNode[int]{Data: 50}, &TreeNode[int]{Data: 200}, &TreeNode[int]{Data: 25}, &TreeNode[int]{Data: 75}, &TreeNode[int]{Data: 300}, &TreeNode[int]{Data: 10}, &TreeNode[int]{Data: 350}, &TreeNode[int]{Data: 15}},
		{&TreeNode[int]{Data: 20}, &TreeNode[int]{Data: 40}, &TreeNode[int]{Data: 50}, &TreeNode[int]{Data: 90}, &TreeNode[int]{Data: 67}, &TreeNode[int]{Data: 94}},
		{&TreeNode[int]{Data: -10}, &TreeNode[int]{Data: -23}, &TreeNode[int]{Data: 45}, &TreeNode[int]{Data: 25}, &TreeNode[int]{Data: 46}},
		{&TreeNode[int]{Data: 9}, &TreeNode[int]{Data: 7}, nil, nil, &TreeNode[int]{Data: 1}, &TreeNode[int]{Data: 8}, &TreeNode[int]{Data: 10}, nil, &TreeNode[int]{Data: 12}},
		{&TreeNode[int]{Data: 3}, &TreeNode[int]{Data: 2}, &TreeNode[int]{Data: 3}, nil, &TreeNode[int]{Data: 3}, nil, &TreeNode[int]{Data: 1}},
	}

	// Create the binary trees using the BinaryTree class
	inputTrees := make([]*BinaryTree[int], len(listOfTrees))
	for i, listOfNodes := range listOfTrees {
		inputTrees[i] = NewBinaryTree[int](listOfNodes)
	}

	// Print the input trees
	for i, tree := range inputTrees {
		fmt.Printf("%d.\tInput Tree:\n", i+1)
		displayTree(tree.root)
		fmt.Printf("\n\tVertical order: %s", arrayToString(verticalOrder(tree.root)))
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
