package main

import "fmt"

type RecentCounter struct {
	sw []int
}

func Constructor() RecentCounter {
	return RecentCounter{sw: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.sw = append(this.sw, t)
	for len(this.sw) > 0 && this.sw[0] < t-3000 {
		this.sw = this.sw[1:]
	}

	return len(this.sw)
}

func main() {
	rc := Constructor()
	fmt.Println("rc.Ping(1),", rc.Ping(1))
	fmt.Println("rc.Ping(100),", rc.Ping(100))
	fmt.Println("rc.Ping(3001),", rc.Ping(3001))
	fmt.Println("rc.Ping(3002),", rc.Ping(3002))

}
