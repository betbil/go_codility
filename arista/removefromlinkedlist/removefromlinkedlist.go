package main

type node struct {
	val  int
	next *node
}

func (n *node) add(r int) *node {
	if n == nil {
		return &node{val: r}
	}
	n.next = n.next.addtolist(r)
	return n
}

func addtolist(list *node, r int) *node {
	if list == nil {
		return &node{val: r}
	}
	list.next = addtolist(list.next, r)
	return list
}

func removematches(list *node, r int) *node {
	if list == nil {
		return nil
	}
	list.next = removematches(list.next, r)
	if list.val == r {
		return list.next
	}
	return list
}

func main() {

}
