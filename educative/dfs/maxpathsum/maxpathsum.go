package main

import (
	"fmt"
	"math"
)

type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

var maxSum = -1 << 31

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxPathSum(root *TreeNode[int]) int {
	// Your code will replace this placeholder return statement
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Data
	}

	leftSum := maxPathSum(root.Left)
	rightSum := maxPathSum(root.Right)

	maxSum = max(maxSum, max(max(leftSum, rightSum), leftSum+rightSum)+root.Data)

	return maxSum
}

// edu
var maxSumedu int = math.MinInt32

// maxPathSum is our solution function //todobetul güzel çözüm
func edumaxPathSum(root *TreeNode[int]) int {
	maxContrib(root)
	temp := maxSumedu
	maxSumedu = math.MinInt32
	return temp
}

// maxContrib is a helper solution function
func maxContrib(root *TreeNode[int]) int {
	if root == nil {
		return 0
	}

	// Sum of the left and right subtree
	maxLeft := maxContrib(root.Left)
	maxRight := maxContrib(root.Right)

	leftSubtree, rightSubtree := 0, 0

	// max sum on the left and right sub-trees of root
	if maxLeft > 0 {
		leftSubtree = maxLeft
	}
	if maxRight > 0 {
		rightSubtree = maxRight
	}

	// the price to start a new path where `root` is a highest root
	priceNewpath := root.Data + leftSubtree + rightSubtree

	// update maxSum if it's better to start a new path
	maxSumedu = max(maxSumedu, priceNewpath)

	// for recursion :
	// return the max contribution if continue the same path
	return root.Data + max(leftSubtree, rightSubtree)
}

//edu

func main() {
	fmt.Println("max path sum", maxPathSum(nil)) // Output: 0

	root := &TreeNode[int]{
		Data: 1,
		Left: &TreeNode[int]{
			Data: 2,
		},
		Right: &TreeNode[int]{
			Data: -1,
		},
	}
	fmt.Println("max path sum", maxPathSum(root)) // Output: 6
	root = &TreeNode[int]{
		Data: -10,
		Left: &TreeNode[int]{
			Data: 9,
			Left: &TreeNode[int]{
				Data: 12,
			},
			Right: &TreeNode[int]{
				Data: 11,
			},
		},
		Right: &TreeNode[int]{
			Data: 20,
			Left: &TreeNode[int]{
				Data: 15,
			},
			Right: &TreeNode[int]{
				Data: 7,
			},
		},
	}

	fmt.Println("max path sum", maxPathSum(root))        // Output: 6
	fmt.Println("edu max path sum", edumaxPathSum(root)) // Output: 6
	//[-10, 9, 20, 12, 11, 15, 7]
}
