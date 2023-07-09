package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    int
	orgindex int
	index    int
}

type maxheap []*Item

func (h *maxheap) Len() int {
	return len(*h)
}

func (h *maxheap) Less(i, j int) bool {
	return (*h)[i].value > (*h)[j].value
}

func (h *maxheap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	(*h)[i].index = i
	(*h)[j].index = j
}

func (h *maxheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *maxheap) Push(x interface{}) {
	n := len(*h)
	item := x.(*Item)
	item.index = n
	*h = append(*h, item)
}

func (h *maxheap) update(item *Item, value int) {
	item.value = value
	heap.Fix(h, item.index)
}

/*
Push(x any) // add x as element Len()

	Pop() any   // remove and return element Len() - 1.
*/
func findMaxSlidingWindow(nums []int, w int) []int {
	fmt.Println("nums", nums, "w", w)
	result := make([]int, 0)
	tmp := make([]*Item, w)
	mxheap := maxheap{}
	heap.Init(&mxheap)

	if len(nums) == 0 || w == 0 {
		return []int{}
	}

	for i := 0; i < w && i < len(nums); i++ {
		tmp[i] = &Item{value: nums[i], orgindex: i}
		heap.Push(&mxheap, tmp[i])
	}
	result = append(result, mxheap[0].value)

	for i := w; i < len(nums); i++ {
		mxheap.update(tmp[i%w], nums[i])
		result = append(result, mxheap[0].value)
	}

	// Your code will replace this placeholder return statement
	return result
}

func main() {
	fmt.Println(findMaxSlidingWindow([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3))
	fmt.Println(findMaxSlidingWindow([]int{1, 2, 3, 4, 3, 2, 1, 2, 5}, 4))
	fmt.Println(findMaxSlidingWindow([]int{10, 6, 9, -3, 23, -1, 34, 56, 67, -1, -4, -8, -2, 9, 10, 34, 67}, 3))
}
