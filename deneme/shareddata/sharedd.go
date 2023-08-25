package main

import (
	"fmt"
	"sync"
	"time"
)

type Shared struct {
	value int
	cond  *sync.Cond
}

func consumer(id int, shared *Shared) {
	shared.cond.L.Lock()
	shared.cond.Wait()
	fmt.Printf("Consumer %d received value: %d\n", id, shared.value)
	shared.cond.L.Unlock()
}

func producer(shared *Shared, value int) {
	shared.cond.L.Lock()
	shared.value = value
	shared.cond.Broadcast() // Notify all waiting consumers
	shared.cond.L.Unlock()
}

func main() {
	var shared Shared
	shared.cond = sync.NewCond(&sync.Mutex{})

	for i := 0; i < 5; i++ {
		go consumer(i, &shared)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Producer is sending value 42")
	producer(&shared, 42)

	time.Sleep(1 * time.Second) // Give consumers time to process
}
