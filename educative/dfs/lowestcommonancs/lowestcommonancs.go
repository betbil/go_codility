package main

type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type Solution struct {
	lca *TreeNode[int]
}

func (s *Solution) lowestCommonAncestor(root, p, q *TreeNode[int]) *TreeNode[int] {
	// Replace this placeholder return statement with your code
	return nil
}

func levelordertraversal(root, p, q *TreeNode[int], tlist *[]int) (int, int) {
	// Your code here
	queue := make([]*TreeNode[int], 0)
	queue = append(queue, root)
	depth := 0
	for len(queue) > 0 {
		depth++
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Data == p.Data {
				*tlist = append(*tlist, depth)
			}
			if node.Data == q.Data {
				*tlist = append(*tlist, depth)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
	}
	return 0, 0
}
