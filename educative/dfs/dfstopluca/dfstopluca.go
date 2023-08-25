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

// todobetul arista
func kthSmallestElementIterative(root *TreeNode[int], k int) int {
	// Replace this placeholder return statement with your code
	stack := make([]*TreeNode[int], 0)

	crnode := root
	for crnode != nil || len(stack) > 0 {
		for crnode != nil {
			stack = append(stack, crnode)
			crnode = crnode.Left
		}
		crnode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return crnode.Data
		}
		crnode = crnode.Right
	}
	return -1
}

func validateBst(root *TreeNode[int]) bool {
	list := make([]int, 0)
	return validateBstRecur(root, &list)
}

func validateBstRecur(root *TreeNode[int], list *[]int) bool {
	if root == nil {
		return true
	}

	if !validateBstRecur(root.Left, list) {
		return false
	}
	*list = append(*list, root.Data)
	if len(*list) > 1 && (*list)[len(*list)-1] < (*list)[len(*list)-2] {
		return false
	}
	if !validateBstRecur(root.Right, list) {
		return false
	}
	return true
}

func maxDepth(root *TreeNode[int]) int {
	return maxDepthrecur(root, 0)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxDepthrecur(root *TreeNode[int], dept int) int {
	if root == nil {
		return dept
	}
	dept++
	return max(maxDepthrecur(root.Left, dept), maxDepthrecur(root.Right, dept))
}

func kthSmallestElement(root *TreeNode[int], k int) int {
	// Replace this placeholder return statement with your code
	list := make([]int, 0, k)
	kthSmallestElementrecur(root, k, &list)
	return list[k-1]
}
func kthSmallestElementrecur(root *TreeNode[int], k int, list *[]int) {
	// Replace this placeholder return statement with your code
	if root == nil {
		return
	}
	kthSmallestElementrecur(root.Left, k, list)
	if len(*list) == k {
		return
	}
	*list = append(*list, root.Data)
	kthSmallestElementrecur(root.Right, k, list)
	return
}

func main() {
	//[2, 1, 3] , 2
	root := &TreeNode[int]{
		Data: 2,
		Left: &TreeNode[int]{
			Data: 1,
		},
		Right: &TreeNode[int]{
			Data: 3,
		},
	}
	fmt.Println(kthSmallestElement(root, 2))          // Output: 2
	fmt.Println(kthSmallestElementIterative(root, 2)) // Output: 2
	fmt.Println(maxDepth(root))                       // Output: 2
	// [5, 3, 6, 2, 4, null, null, 1], 3
	root2 := &TreeNode[int]{
		Data: 5,
		Left: &TreeNode[int]{
			Data: 3,
			Left: &TreeNode[int]{
				Data: 2,
				Left: &TreeNode[int]{
					Data: 1,
				},
			},
			Right: &TreeNode[int]{
				Data: 4,
			},
		},
		Right: &TreeNode[int]{
			Data: 6,
		},
	}
	fmt.Println(kthSmallestElement(root2, 3))          // Output: 3
	fmt.Println(kthSmallestElementIterative(root2, 3)) // Output: 3
	fmt.Println(maxDepth(root2))                       // Output: 4
}
