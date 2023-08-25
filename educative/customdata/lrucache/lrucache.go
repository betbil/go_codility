package main

import "fmt"

// Tip: You may use some of the code templates provided
// in the support files

// We will use a linkedlist of a pair of integers
// where the first integer will be the key
// and the second integer will be the value
type node struct {
	// Write your code
	key   int
	value int
	next  *node
	prev  *node
}

type LRUCache struct {
	// Write your code
	lrumap   map[int]*node
	head     *node
	tail     *node
	capacity int
}

func LRUCacheInitWithCapacity(capacity int) *LRUCache {
	// Replace this placeholder return statement with your code
	lruc := &LRUCache{}
	lruc.lrumap = make(map[int]*node)
	lruc.capacity = capacity
	lruc.head = nil
	lruc.tail = nil
	return lruc
}

func (l *LRUCache) Get(key int) int {
	// Replace this placeholder return statement with your code
	val, ok := l.lrumap[key]
	if ok {
		l.updateHeadTail(val, val.key == l.tail.key)
		return val.value
	}
	return -1
}

func (l *LRUCache) updateHeadTail(newn *node, updateTail bool) {
	if l.head == nil {
		l.head = newn
		l.tail = newn
	} else {
		//
		if newn.prev != nil {
			newn.prev.next = newn.next
		}
		if newn.next != nil {
			newn.next.prev = newn.prev
		}
		newn.next = l.head

		l.head.prev = newn
		l.head = newn
		if updateTail {
			l.tail = l.tail.prev
			l.tail.next = nil
		}
		newn.prev = nil
	}
}

func (l *LRUCache) Set(key int, value int) {
	// Write your code here
	val, ok := l.lrumap[key]
	if ok {
		// update the value
		val.value = value
		l.updateHeadTail(val, val.key == l.tail.key)
	} else {
		nodev := &node{key: key, value: value}
		updateTail := false
		if len(l.lrumap) == l.capacity {
			// remove the tail
			delete(l.lrumap, l.tail.key)
			updateTail = true
		}
		l.lrumap[key] = nodev
		l.updateHeadTail(nodev, updateTail)
	}
}

func main() {
	//["LRUCache", "Get", "Set", "Get", "Set", "Get", "Set", "Get", "Get"] , [[2], [10], [10, 100], [10], [15, 50], [10], [30, 300], [15], [30]]
	/*
		l := LRUCacheInitWithCapacity(2)
		fmt.Println(l.Get(10))
		l.Set(10, 100)
		fmt.Println(l.Get(10))
		l.Set(15, 50)
		fmt.Println(l.Get(10))
		l.Set(30, 300)
		fmt.Println(l.Get(15))
		fmt.Println(l.Get(30))


	*/

	//["LRUCache", "Set", "Set", "Get", "Set", "Set", "Set", "Get", "Set", "Get"] , [[3], [50, 50], [51, 51], [51], [52, 52], [53, 53], [54, 54], [53], [55, 55], [51]]
	l := LRUCacheInitWithCapacity(3)
	l.Set(50, 50)
	l.Set(51, 51)
	fmt.Println(l.Get(51))
	l.Set(52, 52)
	l.Set(53, 53)
	l.Set(54, 54)
	fmt.Println(l.Get(53))
	l.Set(55, 55)
	fmt.Println(l.Get(51))
}
