package main

import "container/list"

// Tip: You may use some of the code templates provided
// in the support files

// We will use a linkedlist of a pair of integers
// where the first integer will be the key
// and the second integer will be the value

type pair struct {
	value int
	key   int
}
type LRUCache2 struct {
	// Write your code
	lrumap     map[int]*list.Element
	linkedlist *list.List
	capacity   int
}

func LRUCacheInitWithCapacity2(capacity int) *LRUCache2 {
	// Replace this placeholder return statement with your code
	l := &LRUCache2{}
	l.lrumap = make(map[int]*list.Element)
	l.linkedlist = list.New()
	l.capacity = capacity
	return l
}

func (l *LRUCache2) Get(key int) int {
	// Replace this placeholder return statement with your code
	val, ok := l.lrumap[key]
	if ok {
		l.linkedlist.MoveToFront(val)
		return val.Value.(pair).value
	}
	return -1
}

func (l *LRUCache2) Set(key int, value int) {
	// Write your code here
	val, ok := l.lrumap[key]
	if ok {
		val.Value = pair{value, key}
		l.linkedlist.MoveToFront(val)
		return
	}

	if len(l.lrumap) == l.capacity {
		e := l.linkedlist.Back()
		delete(l.lrumap, e.Value.(pair).key)
		l.linkedlist.Remove(e)
	}
	e := l.linkedlist.PushFront(pair{value, key})
	l.lrumap[key] = e
}
