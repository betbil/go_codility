package main

import "fmt"

// todobetul check
func detectCycle(head *EduLinkedListNode) bool {
	// Your code will replce this placeholder return statement

	slow := head
	fast := head

	for fast != nil {
		slow = slow.next
		if fast.next == nil {
			return false
		}
		fast = fast.next.next

		if slow == fast {
			return true
		}

	}
	return false
}

func main() {
	// Test your function
	// Create a linked list
	llist := new(EduLinkedList)
	llist.CreateLinkedList([]int{1, 2, 3, 4, 5, 6, 4})
	fmt.Println("detectCycle", detectCycle(llist.head))
}
