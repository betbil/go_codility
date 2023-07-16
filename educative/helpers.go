package educative

import "strconv"

func stackhelper() {
	stack := make([]int, 0, 5)
	nums2 := []int{1, 2, 3, 4, 5}
	m := make(map[int]int)
	// iterate over nums2
	for _, current := range nums2 {
		// while stack is not empty and current element is greater than the top element of the stack
		for len(stack) > 0 && current > stack[len(stack)-1] {
			// update the map with the current element as the value for the popped element
			m[stack[len(stack)-1]] = current
			// pop the top element from the stack
			stack = stack[:len(stack)-1]
		}
		// push the current element to the stack
		stack = append(stack, current)
	}
}

type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func queuehelper(root *TreeNode[int]) string {
	queue := make([]*TreeNode[int], 0)
	queue = append(queue, root)
	depth := 0
	outstr := make([]byte, 0) //rune da olabilir
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
