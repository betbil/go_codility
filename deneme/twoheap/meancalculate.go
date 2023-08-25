package main

import (
	"container/heap"
	"fmt"
)

// Min heap
type MinHeap []float64

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Max heap
type MaxHeap []float64

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type RunningStat struct {
	mins MinHeap
	maxs MaxHeap
	sum  float64
}

func (rs *RunningStat) Add(num float64) {
	// Add to one of the heaps
	if rs.maxs.Len() == 0 || num <= rs.maxs[0] {
		heap.Push(&rs.maxs, num)
	} else {
		heap.Push(&rs.mins, num)
	}

	// Balance heaps
	if rs.maxs.Len() > rs.mins.Len()+1 {
		heap.Push(&rs.mins, heap.Pop(&rs.maxs))
	} else if rs.mins.Len() > rs.maxs.Len() {
		heap.Push(&rs.maxs, heap.Pop(&rs.mins))
	}

	rs.sum += num
}

func (rs *RunningStat) Mean() float64 {
	return rs.sum / float64(rs.mins.Len()+rs.maxs.Len())
}

func (rs *RunningStat) Median() float64 {
	if rs.mins.Len() > rs.maxs.Len() {
		return rs.mins[0]
	} else if rs.maxs.Len() > rs.mins.Len() {
		return rs.maxs[0]
	}
	return (rs.mins[0] + rs.maxs[0]) / 2
}

func main() {
	numbers := []float64{12, 4, 5, 3, 8, 7}
	rs := &RunningStat{}
	for _, num := range numbers {
		rs.Add(num)
		fmt.Printf("Added %v. Mean: %v, Median: %v\n", num, rs.Mean(), rs.Median())
	}
}
