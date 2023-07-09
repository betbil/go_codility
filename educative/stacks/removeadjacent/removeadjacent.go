package main

import (
	"container/list"
)

func removeDuplicates(toCleanUp string) string {

	// Write your code here
	if len(toCleanUp) < 2 {
		return toCleanUp
	}

	stack := list.New()

	for _, c := range toCleanUp {
		if stack.Len() == 0 {
			stack.PushBack(c)
		} else {
			if stack.Back().Value != c {
				stack.PushBack(c)
			} else {
				stack.Remove(stack.Back())
			}
		}
	}

	bytes := make([]byte, stack.Len())
	for i := stack.Len() - 1; i >= 0; i-- {
		bytes[i] = byte(stack.Remove(stack.Back()).(rune))
	}

	return string(bytes)

}
