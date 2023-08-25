package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func (bst *BST) Insert(val int) {
	/*
		if bst.Root == nil {
			bst.Root = &Node{Val: val}
			return
		}
	*/
	//bst.Root = insertRec(bst.Root, val)
	fmt.Printf("bst.Root (*Node) p:%p: \n", bst.Root)
	fmt.Printf("&(bst.Root) (**Node) p:%p: \n", &(bst.Root))
	insertRec2(&(bst.Root), val)
}

func insertRec2(node **Node, val int) {
	fmt.Printf("**Node p:%p: \n", node)
	if *node == nil {
		*node = &Node{Val: val}
		fmt.Printf("*Node p:%p: \n", *node)
		return
	}
	if val < (*node).Val {

		insertRec2(&((*node).Left), val)
	} else {
		insertRec2(&((*node).Right), val)
	}
	return
}

func insertRec(node *Node, val int) *Node {
	if node == nil {
		return &Node{Val: val}
	}
	if val < node.Val {
		node.Left = insertRec(node.Left, val)
	} else {
		node.Right = insertRec(node.Right, val)
	}
	return node
}

func main() {
	bst := &BST{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(2)
	bst.Insert(5)
	bst.Insert(13)
	bst.Insert(22)
	bst.Insert(1)
	bst.Insert(14)
	bst.Insert(12)
	bst.Insert(11)
	bst.Insert(3)
}
