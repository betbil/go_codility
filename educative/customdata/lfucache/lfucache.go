package main

import "container/heap"

type node struct {
	key        int
	val        int
	frequency  int
	updatetime int
	index      int
}

// Min heap
type MinHeap []*node

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].frequency < h[j].frequency {
		return true
	} else if h[i].frequency == h[j].frequency {
		return h[i].updatetime < h[j].updatetime
	}
	return false
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *MinHeap) Push(x interface{}) {
	n := len(*h)
	*h = append(*h, x.(*node))
	(*h)[n].index = n
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil // avoid memory leak
	x.index = -1   // for safety
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) update(n *node, value, frequency, updatetime int) {
	n.val = value
	n.frequency = frequency
	n.updatetime = updatetime
	heap.Fix(h, n.index)
}

type LFUCache struct {
	// Write your code here
	capacity int
	lfumap   map[int]*node
	minheap  *MinHeap
}

// constructor intializes the LFUCache type object
func Constructor(capacity int) *LFUCache {
	l := LFUCache{}
	l.capacity = capacity
	l.lfumap = make(map[int]*node)
	l.minheap = &MinHeap{}
	// Replace this placeholder return statement with your code
	return &l
}

// Get returns value for the given key
func (this *LFUCache) Get(key int) int {
	// Replace this placeholder return statement with your code
	val, ok := this.lfumap[key]
	if ok {
		this.minheap.update(val, val.val, val.frequency+1, val.updatetime+1)
		return val.val
	}
	return -1
}

// Put sets the value at a given key.
func (this *LFUCache) Put(key int, value int) {
	// Write your code here
	if len(this.lfumap) < this.capacity {
		// insert
		val, ok := this.lfumap[key]
		if ok {
			this.minheap.update(val, value, val.frequency+1, val.updatetime+1)
		} else {
			n := &node{key: key, val: value, frequency: 1, updatetime: 1}
			this.lfumap[key] = n
			heap.Push(this.minheap, n)
		}
	} else {
		// update
		val, ok := this.lfumap[key]
		if ok {
			this.minheap.update(val, value, val.frequency+1, val.updatetime+1)
		} else {
			// remove min
			n := heap.Pop(this.minheap).(*node)
			delete(this.lfumap, n.key)
			n = &node{key: key, val: value, frequency: 1, updatetime: 1}
			this.lfumap[key] = n
			heap.Push(this.minheap, n)
		}
	}
}
