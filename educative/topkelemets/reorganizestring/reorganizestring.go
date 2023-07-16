package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	frq int
	val byte
}

type maxheap []*Item

func (h maxheap) Len() int {
	return len(h)
}

func (h maxheap) Less(i, j int) bool {
	return h[i].frq > h[j].frq
}

func (h maxheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxheap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *maxheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h maxheap) Top() *Item {
	return h[0]
}

func reorganizeString(str string) string {
	fmt.Println("received string:", str)
	// Write your code here
	frqMap := make(map[byte]int)
	for _, char := range str {
		frqMap[byte(char)]++
	}

	mxHeap := &maxheap{}
	heap.Init(mxHeap)
	for char, frq := range frqMap {
		heap.Push(mxHeap, &Item{frq, char})
	}

	strlength := len(str)
	bytes := make([]byte, 0)
	var poppedTmp, popped *Item
	for mxHeap.Len() > 0 {
		if poppedTmp != nil {
			popped = poppedTmp
			poppedTmp = nil
		} else {
			popped = heap.Pop(mxHeap).(*Item)
		}

		if popped.frq > (strlength+1)/2 {
			return ""
		}
		bytes = append(bytes, popped.val)
		for i := 1; i < popped.frq; i++ {
			if poppedTmp == nil {
				poppedTmp = heap.Pop(mxHeap).(*Item)
			}
			bytes = append(bytes, poppedTmp.val)
			poppedTmp.frq--
			if poppedTmp.frq == 0 {
				poppedTmp = nil
			}
			bytes = append(bytes, popped.val)
		}
	}

	return string(bytes)
}

func main() {
	fmt.Println(reorganizeString("aaab"))
	fmt.Println(reorganizeString("aaabccc"))
}
