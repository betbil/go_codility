// Template for Dequeue

package helper

import (
	"strconv"
)

type Dequeue struct {
	items []*TreeNode[int]
}

// NewDequeue is a constructor that will declare and return the dequeue type object
func NewDequeue() *Dequeue {
	return new(Dequeue)
}

// todobetul güzel çözüm
// PushFront will push an element at the front of the dequeue
func (s *Dequeue) PushFront(item *TreeNode[int]) {
	temp := []*TreeNode[int]{item}
	s.items = append(temp, s.items...)
}

// PushBack will push an element at the back of the dequeue
func (s *Dequeue) PushBack(item *TreeNode[int]) {
	s.items = append(s.items, item)
}

// PopFront will pop an element from the front of the dequeue
func (s *Dequeue) PopFront() *TreeNode[int] {
	defer func() {
		s.items = s.items[1:]
	}()
	return s.items[0]
}

// PopBack will pop an element from the back of the dequeue
func (s *Dequeue) PopBack() *TreeNode[int] {
	i := len(s.items) - 1
	defer func() {
		s.items = append(s.items[:i], s.items[i+1:]...)
	}()
	return s.items[i]
}

// Front will return the element from the front of the dequeue
func (s *Dequeue) Front() *TreeNode[int] {
	return s.items[0]
}

// Back will return the element from the back of the dequeue
func (s *Dequeue) Back() *TreeNode[int] {
	return s.items[len(s.items)-1]
}

// Empty will check if the dequeue is empty or not
func (s *Dequeue) Empty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

// Len will return the length of the dequeue
func (s *Dequeue) Len() int {
	return len(s.items)
}

// String function converts queue elements to string
func (s *Dequeue) String() string {
	output := "["
	for i := range s.items {
		output += strconv.Itoa(s.items[i].Data) + ", "
	}
	return output[:len(output)-2] + "]"
}
