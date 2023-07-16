package main

import (
	"container/heap"
	"fmt"
)

// Tip: You may use some of the code templates provided
// in the support files

type heapst struct {
	values []int
	less   func(int, int) bool
}

func (h *heapst) Len() int {
	return len(h.values)
}

func (h *heapst) Less(i, j int) bool {
	return h.less(h.values[i], h.values[j])
}

func (h *heapst) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *heapst) Push(x interface{}) {
	h.values = append(h.values, x.(int))
}

func (h *heapst) Pop() interface{} {
	old := h.values
	n := len(old)
	x := old[n-1]
	h.values = old[0 : n-1]
	return x
}

type MedianOfStream struct {
	// Write your code here
	minheap *heapst
	maxheap *heapst
	isInit  bool
}

func (this *MedianOfStream) Init() {
	// Write your code here
	fmt.Println("Inıtttt")
	if !this.isInit {
		this.minheap = &heapst{values: make([]int, 0), less: func(i, j int) bool { return i < j }}
		this.maxheap = &heapst{values: make([]int, 0), less: func(i, j int) bool { return i > j }}
		heap.Init(this.minheap)
		heap.Init(this.maxheap)
		this.isInit = true
	}

}

// insertNum() should take a number and store it
func (this *MedianOfStream) InsertNum(num int) float64 {
	this.Init()
	fmt.Println("InsertNum this.minheap.Len()", this.minheap.Len(), "this.maxheap.Len()", this.maxheap.Len())
	// Your code will replace this placeholder return statement
	if this.minheap.Len() == 0 || num > this.minheap.values[0] {
		heap.Push(this.minheap, num)
	} else {
		heap.Push(this.maxheap, num)
	}

	//balance
	if this.minheap.Len() > this.maxheap.Len()+1 {
		heap.Push(this.maxheap, heap.Pop(this.minheap))
	} else if this.maxheap.Len() > this.minheap.Len()+1 {
		heap.Push(this.minheap, heap.Pop(this.maxheap))
	}

	return -99.99
}

// todobetul check
// findMedian() should return the median of the stored numbers
func (this *MedianOfStream) FindMedian() float64 {
	this.Init()
	// Your code will replace this placeholder return statement
	fmt.Println("FindMedian this.minheap.Len()", this.minheap.Len(), "this.maxheap.Len()", this.maxheap.Len())
	if this.minheap.Len() > this.maxheap.Len() {
		return float64(this.minheap.values[0])
	} else if this.maxheap.Len() > this.minheap.Len() {
		return float64(this.maxheap.values[0])
	} else if this.maxheap.Len() == this.minheap.Len() && this.maxheap.Len() > 0 {
		return float64(this.minheap.values[0]+this.maxheap.values[0]) / 2.0
	}
	return -99.99
}

func medianSlidingWindow(nums []int, k int) []float64 {
	// Your code will replace this placeholder return statement
	outgoing := make(map[int]int)

	smalls := &heapst{values: make([]int, 0), less: func(i, j int) bool { return i > j }} //maxheap of mins
	larges := &heapst{values: make([]int, 0), less: func(i, j int) bool { return i < j }} //minheap of maxs
	heap.Init(smalls)
	heap.Init(larges)

	medians := make([]float64, 0)

	for i, num := range nums {
		if i >= k {
			outgoing[nums[i-k]]++
			for k, v := range outgoing {
				if k == smalls.values[0] || k == larges.values[0] {
					if k == smalls.values[0] {
						heap.Pop(smalls)
					} else {
						heap.Pop(larges)
					}
					outgoing[k]--
					if v == 1 {
						delete(outgoing, k)
					}
				}
			}
		}

		if smalls.Len() == 0 || num < smalls.values[0] {
			heap.Push(smalls, num)
		} else {
			heap.Push(larges, num)
		}

		//balance
		if smalls.Len() > larges.Len()+1 {
			heap.Push(larges, heap.Pop(smalls))
		} else if larges.Len() > smalls.Len()+1 {
			heap.Push(smalls, heap.Pop(larges))
		}
		if i >= k-1 {
			if k%2 == 1 {
				medians = append(medians, float64(smalls.values[0]))
			} else {
				medians = append(medians, float64(smalls.values[0]+larges.values[0])/2.0)
			}
		}
	}

	return medians
}

func main() {
	median := MedianOfStream{}
	fmt.Println("The median is: ", median.InsertNum(3))
	fmt.Println("The median is: ", median.InsertNum(1))
	fmt.Println("The median is: ", median.FindMedian())
	fmt.Println("The median is: ", median.InsertNum(5))
	fmt.Println("The median is: ", median.FindMedian())
	fmt.Println("The median is: ", median.InsertNum(4))
	fmt.Println("The median is: ", median.FindMedian())

	fmt.Println("medianSlidingWindow")
	fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))

}
