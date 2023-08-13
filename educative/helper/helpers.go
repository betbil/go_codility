package helper

import (
	"math"
	"strconv"
)

func reverse(sw []int) {
	for a, b := 0, len(sw)-1; a < b; a, b = a+1, b-1 {
		sw[a], sw[b] = sw[b], sw[a]
	}
}

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

//todobetul heapleride al

// set implementation
type Set[T comparable] struct {
	m map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]bool)}
}

func sethelper() {
	nums := []int{100, 4, 200, 1, 3, 2}
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
}

// bitset helpers
// pow calculates the power of the given digit
func pow(digit int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * digit
	}
	return res
}

func findBitwiseComplement(num int) int {
	if num == 0 { // if the value of num is 0, return 1
		return 1
	}

	// converting the value into its binary representation and
	// counting the number of bits required by this number
	bitCount := math.Floor(math.Log2(float64(num))) + 1
	// computing the all bits set of the number
	allBitsSet := pow(2, int(bitCount)) - 1

	// flipping all bits of number by taking xor with all_bits_set
	return num ^ allBitsSet
}

// heap 1
type heapst [][]int

func (h heapst) Len() int {
	return len(h)
}

func (h heapst) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h heapst) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapst) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *heapst) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h heapst) Top() int {
	return h[0][0]
}

// heap 1
// heap 2
type minheap []int

func (h minheap) Len() int {
	return len(h)
}

func (h minheap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h minheap) Top() int {
	return h[0]
}

//heap 2

//heap 3

type Item struct {
	value    int
	orgindex int
}
type minheap []*Item
type maxheap []*Item

func (h minheap) Len() int {
	return len(h)
}

func (h minheap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h minheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Item))
}

func (h *minheap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *minheap) Top() *Item {
	return (*h)[0]
}

func (h maxheap) Len() int {
	return len(h)
}

func (h maxheap) Less(i, j int) bool {
	return h[i].value > h[j].value
}

func (h maxheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxheap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Item))
}

func (h *maxheap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//heap 3
