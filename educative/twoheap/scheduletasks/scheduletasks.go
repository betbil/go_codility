package main

import "container/heap"

type heapst [][]int //start vs end || end vs mcnum

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

func (h heapst) Top() []int {
	return h[0]
}

func tasks(T [][]int) int {
	// Write your code here
	// Your code will replace this placeholder return statement
	startTimeHeap := &heapst{}
	endTimeHeap := &heapst{}
	heap.Init(startTimeHeap)
	heap.Init(endTimeHeap)
	for _, t := range T {
		heap.Push(startTimeHeap, t)
	}

	mcnum := 0
	for startTimeHeap.Len() > 0 {
		tmpEarly := heap.Pop(startTimeHeap).([]int)
		if endTimeHeap.Len() > 0 && endTimeHeap.Top()[0] <= tmpEarly[0] {
			popped := heap.Pop(endTimeHeap).([]int)
			heap.Push(endTimeHeap, []int{tmpEarly[1], popped[1]})
		} else {
			mcnum++
			heap.Push(endTimeHeap, []int{tmpEarly[1], mcnum})
		}
	}

	return mcnum
}
