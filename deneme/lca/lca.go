package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive function to find the LCA of two nodes n1 and n2.
func findLCA(root, n1, n2 *TreeNode) *TreeNode {
	var fLCAImpl func(root *TreeNode) *TreeNode
	fLCAImpl = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		if root.Value == n1.Value || root.Value == n2.Value {
			return root
		}

		leftLCA := fLCAImpl(root.Left)
		rightLCA := fLCAImpl(root.Right)

		if leftLCA != nil && rightLCA != nil {
			return root
		}

		if leftLCA != nil {
			return leftLCA
		}
		return rightLCA
	}

	return fLCAImpl(root)
}

func main() {
	// Constructing a sample binary tree:
	//         3
	//       /   \
	//      5     1
	//     / \   / \
	//    6   2 0   8
	//       / \
	//      7   4
	root := &TreeNode{3, nil, nil}
	root.Left = &TreeNode{5, nil, nil}
	root.Right = &TreeNode{1, nil, nil}
	root.Left.Left = &TreeNode{6, nil, nil}
	root.Left.Right = &TreeNode{2, nil, nil}
	root.Left.Right.Left = &TreeNode{7, nil, nil}
	root.Left.Right.Right = &TreeNode{4, nil, nil}
	root.Right.Left = &TreeNode{0, nil, nil}
	root.Right.Right = &TreeNode{8, nil, nil}

	// Example: LCA of nodes 5 and 1 is 3.
	n1 := root.Left  // Node with value 5.
	n2 := root.Right // Node with value 1.
	lca := findLCA(root, n1, n2)
	fmt.Printf("LCA of %d and %d is %d\n", n1.Value, n2.Value, lca.Value)
	n1 = root.Left.Right.Left  // Node with value 5.
	n2 = root.Left.Right.Right // Node with value 1.
	lca = findLCA(root, n1, n2)
	fmt.Printf("LCA of %d and %d is %d\n", n1.Value, n2.Value, lca.Value)
}
