package main

import (
	"container/heap"
	"fmt"
)

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

func kSmallestNumber(lists [][]int, k int) int {
	// Write your code here
	// Your code will replace this placeholder return statement
	minheap := &heapst{}
	heap.Init(minheap)

	for i, list := range lists {
		if len(list) == 0 {
			continue
		}
		heap.Push(minheap, []int{list[0], i})
	}

	for i := 1; i <= k && minheap.Len() > 0; i++ {
		tmp := heap.Pop(minheap).([]int)
		if i == k {
			return tmp[0]
		}
		//todobetul check güzel
		if len(lists[tmp[1]]) > 1 {
			heap.Push(minheap, []int{lists[tmp[1]][1], tmp[1]})
			lists[tmp[1]] = lists[tmp[1]][1:]
		}
		if minheap.Len() == 0 {
			return tmp[0]
		}
	}

	return 0
}

func main() {
	fmt.Println(kSmallestNumber([][]int{{2, 6, 8}, {3, 6, 7}, {1, 3, 4}}, 5))
	fmt.Println(kSmallestNumber([][]int{{}, {}, {}}, 5))
}
