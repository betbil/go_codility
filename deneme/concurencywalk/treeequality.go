package main

import (
	"bytes"
	"fmt"
	"golang.org/x/tour/tree"
	"runtime"
	"time"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, stop chan bool) {
	defer close(ch)
	var valki func(t *tree.Tree)
	valki = func(t *tree.Tree) {
		if t == nil {
			return
		}

		valki(t.Left)
		select {
		case ch <- t.Value:
			fmt.Println(t.Value, " send")
			// Data sent successfully.
			/*
				case <-stop:
						fmt.Println("Channel is closed, not sending data.")
						return
			*/

		case val, ok := <-stop:
			if !ok {
				fmt.Println("Channel is closed, not sending data.")
				return
			} else {
				fmt.Println("val", val)
			}

		}
		valki(t.Right)
	}
	valki(t)

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	stop := make(chan bool)
	go Walk(t1, c1, stop)
	go Walk(t2, c2, stop)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 != ok2 || v1 != v2 {
			stop <- true
			fmt.Println("stop")
			close(stop)
			return false
		}

		if !ok1 && !ok2 {
			break
		}
	}

	return true

}

func main2() {
	fmt.Println("Trees equivalent?", Same(tree.New(5), tree.New(2)))
}

func getGID() uint64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := bytes.Fields(buf[:n])[1]
	var id uint64
	for _, b := range idField {
		id = id*10 + uint64(b-'0')
	}
	return id
}

func threadFunc(ch chan int) {
	fmt.Println("Gorutine ", getGID(), " started")
	select {
	case v, ok := <-ch:
		if !ok {
			fmt.Println("Channel is closed, in gorutine ", getGID())
		} else {
			fmt.Println(v, " received in gorutine ", getGID())
		}
	}
	return
}
func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go threadFunc(ch)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("sending 1 to channel")
	ch <- 1
	time.Sleep(10 * time.Second)
	close(ch)
	time.Sleep(5 * time.Second)
}
