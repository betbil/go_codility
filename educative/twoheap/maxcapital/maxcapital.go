package main

import (
	"container/heap"
	"fmt"
)

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

func (h *maxheap) Top() *Item {
	return (*h)[0]
}

// todobetul check
func maximumCapital(c int, k int, capitals []int, profits []int) int {
	// your code will replace the placeholder return statement below
	cpheap := minheap{}
	heap.Init(&cpheap)

	for i, cp := range capitals {
		heap.Push(&cpheap, &Item{cp, i})
	}

	totalCap := c
	for k > 0 {
		pheap := maxheap{}
		heap.Init(&pheap)
		poplist := []*Item{}
		i := 0
		for cpheap.Len() > 0 {
			poped := heap.Pop(&cpheap).(*Item)
			poplist = append(poplist, poped)
			if poped.value <= totalCap {
				heap.Push(&pheap, &Item{profits[poped.orgindex], i})
				i++
			} else {
				break
			}
		}
		if pheap.Len() == 0 {
			break
		}
		maxprofit := heap.Pop(&pheap).(*Item)
		totalCap += maxprofit.value
		k--
		for i, item := range poplist {
			if i != maxprofit.orgindex {
				heap.Push(&cpheap, item)
			}
		}
	}

	return totalCap
}

func main() {
	fmt.Println(maximumCapital(0, 2, []int{0, 1, 2}, []int{1, 2, 3}))
}
