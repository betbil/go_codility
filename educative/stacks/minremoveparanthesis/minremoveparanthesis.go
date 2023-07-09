package main

import (
	"container/list"
	"fmt"
)

type StackElement struct {
	indice    int
	character rune
}

func minRemoveParentheses(s string) string {
	// your code will replace this placeholder return statement
	stack := list.New()
	for i, c := range s {
		if c == '(' {
			stack.PushBack(StackElement{i, c})
		} else if c == ')' {
			if stack.Len() > 0 {
				if (stack.Back().Value.(StackElement)).character == '(' {
					stack.Remove(stack.Back())
				} else {
					stack.PushBack(StackElement{i, c})
				}
			} else {
				stack.PushBack(StackElement{i, c})
			}
		}
	}

	if stack.Len() > 0 {
		//bytes := make([]byte, 0)
		var bytes []byte
		element := stack.Front()
		if element == nil {
			return s
		}
		for i, c := range s {
			if element != nil {
				if i != element.Value.(StackElement).indice {
					bytes = append(bytes, byte(c))
				} else {
					element = element.Next()
				}
			} else {
				bytes = append(bytes, byte(c))
			}

		}
		return string(bytes)
	}

	return s
}

func main() {
	fmt.Println(minRemoveParentheses("("))
	fmt.Println(minRemoveParentheses("ab)ca(so)(sc(s)("))
	fmt.Println("finito")
}
