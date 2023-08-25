package main

import "fmt"

// MinStack declaration
type MinStack struct {
	stack   *[]int
	minlist *[]int
}

func NewMinStack() *MinStack {
	// Replace this placeholder return statement with your code
	mstack := MinStack{}
	mstack.stack = new([]int)
	mstack.minlist = new([]int)
	*mstack.stack = make([]int, 0)
	*mstack.minlist = make([]int, 0)
	return &mstack
}

// Pop removes and returns value from stack
func (m *MinStack) Pop() int {
	// Replace this placeholder return statement with your code
	*m.minlist = (*m.minlist)[:len(*m.minlist)-1]
	x := (*m.stack)[len(*m.stack)-1]
	*m.stack = (*m.stack)[:len(*m.stack)-1]
	return x
}

// Push pushes value into the stack
func (m *MinStack) Push(value int) {
	*m.stack = append(*m.stack, value)
	if len(*m.minlist) == 0 || value < (*m.minlist)[len(*m.minlist)-1] {
		*m.minlist = append(*m.minlist, value)
	} else {
		*m.minlist = append(*m.minlist, (*m.minlist)[len(*m.minlist)-1])
	}
}

// MinNumber returns minimum value in O(1)
func (m *MinStack) MinNumber() int {
	return (*m.minlist)[len(*m.minlist)-1]
}

func main() {
	//["MinStack", "Push()", "Push()", "Pop()", "Push()", "Pop()", "MinNumber()"] , [null, 5, 10, null, 23, null, null]
	mstack := NewMinStack()
	mstack.Push(5)
	mstack.Push(10)
	fmt.Println("mstack.Pop()", mstack.Pop())
	mstack.Push(23)
	fmt.Println("mstack.Pop()", mstack.Pop())
	fmt.Println("mstack.MinNumber()", mstack.MinNumber())
}
