package main

import "fmt"

//todobetul check

func removeNthLastNode(head *EduLinkedListNode, n int) *EduLinkedListNode {
	var eduLinkedList = EduLinkedList{head: head}
	fmt.Print("-------- received head list ---------\n	")
	eduLinkedList.DisplayLinkedList()
	fmt.Println("n: ", n)

	f := head
	s := head
	var pre *EduLinkedListNode
	for i := 0; i < n; i++ {
		if f == nil {
			return nil
		}
		fmt.Println("f node: ", f.data)
		f = f.next
	}

	for f != nil {
		fmt.Println("f node: ", f.data, " s node: ", s.data)
		f = f.next
		pre = s
		s = s.next
	}

	//remove node s
	if pre != nil {
		pre.next = s.next
		fmt.Println("pre node: ", pre.data)
		fmt.Println("removed node: ", s.data)
	} else {
		fmt.Println("pre nii-ll: ")
		head = s.next
	}

	fmt.Print("After remove list")
	(&EduLinkedList{head: head}).DisplayLinkedList()
	return head
}

// todobetul check
func eduSolution_removeNthLastNode(head *EduLinkedListNode, n int) *EduLinkedListNode {

	// Point two pointers, right and left, at head.
	right := head
	left := head

	// Move right pointer n elements away from the left pointer.
	for i := 0; i < n; i++ {
		right = right.next
	}

	// Removal of the head node.
	if right == nil {
		return head.next
	}

	// Move both pointers until right pointer reaches the last node.
	for right.next != nil {
		right = right.next
		left = left.next
	}

	// At this point, left pointer points to (n-1)th element.
	// So link it to next to next element of left.
	left.next = left.next.next

	return head
}

func reverse(head *EduLinkedListNode) *EduLinkedListNode {
	// Write your code here
	// Your code will replace this placeholder return statement
	var pre, next *EduLinkedListNode
	cur := head
	for cur != nil {
		next = cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	fmt.Println("---")
	fmt.Print("After reversing list")
	(&EduLinkedList{head: pre}).DisplayLinkedList()
	return pre
}

func reverseInternalKGroups(head *EduLinkedListNode, k int) (*EduLinkedListNode, *EduLinkedListNode) {
	// Write your code here
	// Your code will replace this placeholder return statement
	var pre, next *EduLinkedListNode
	cur := head
	for cur != nil && k > 0 {
		next = cur.next
		cur.next = pre
		pre = cur
		cur = next
		k--
	}
	return pre, head
}

// todobetul check
func reverseKGroups(head *EduLinkedListNode, k int) *EduLinkedListNode {
	// Write your code here
	// Your code will replace this placeholder return statement
	var before *EduLinkedListNode
	cur := head
	groupsize := k
	for cur != nil {
		headtmp := cur
		for cur != nil && k > 0 {
			cur = cur.next
			k--
		}

		if k == 0 {
			k = groupsize
			beforeNew, afterNew := reverseInternalKGroups(headtmp, k)
			afterNew.next = cur
			if before != nil {
				before.next = beforeNew
			} else {
				head = beforeNew
			}
			before = afterNew
		}
	}

	fmt.Println("---")
	fmt.Print("After reversing K group list")
	(&EduLinkedList{head: head}).DisplayLinkedList()
	return head
}

func main() {
	// Test your function here
	var eduLinkedList = EduLinkedList{}
	eduLinkedList.CreateLinkedList([]int{1, 2, 3, 4, 5})

	//fmt.Println(removeNthLastNode(&EduLinkedListNode{1, &EduLinkedListNode{2, &EduLinkedListNode{3, &EduLinkedListNode{4, &EduLinkedListNode{5, nil}}}}}, 2))
	removeNthLastNode(eduLinkedList.head, 2)
	eduLinkedList2 := EduLinkedList{}
	eduLinkedList2.CreateLinkedList([]int{69, 8, 49, 106, 116, 112})
	removeNthLastNode(eduLinkedList2.head, 6)
	eduLinkedList3 := EduLinkedList{}
	eduLinkedList3.CreateLinkedList([]int{69, 8, 49, 106, 116, 112})
	reverse(eduLinkedList3.head)
	eduLinkedList4 := EduLinkedList{}
	eduLinkedList4.CreateLinkedList([]int{69, 8, 49, 106, 116, 112})
	reverseKGroups(eduLinkedList4.head, 3)

}
