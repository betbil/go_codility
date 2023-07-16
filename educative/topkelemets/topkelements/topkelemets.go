package main

import "container/heap"

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

type KthLargest struct {
	// Write your code here
	mheap *minheap
	kVal  int
}

// KthLargestInit is a constructor to initialize heap and store values in it
func (this *KthLargest) KthLargestInit(k int, nums []int) {
	// Write your code here
	this.mheap = &minheap{}
	this.kVal = k
	heap.Init(this.mheap)

	for i, num := range nums {
		if i < this.kVal {
			heap.Push(this.mheap, num)
			continue
		}

		if num > this.mheap.Top() {
			heap.Pop(this.mheap)
			heap.Push(this.mheap, num)
		}
	}
}

// Add function adds element in the heap
func (this *KthLargest) Add(value int) int {
	// Write your code here
	// Your code will replace this placeholder return statement
	if this.mheap.Len() < this.kVal {
		heap.Push(this.mheap, value)
		return this.mheap.Top()
	} else {
		if value > this.mheap.Top() {
			heap.Pop(this.mheap)
			heap.Push(this.mheap, value)
			return value
		}
	}
	return -1
}

// Returns the Kth largest element from the heap
func (this *KthLargest) ReturnKthLargest() int {
	// Write your code here
	// Your code will replace this placeholder return statement
	return this.mheap.Top()
}
