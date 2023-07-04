package main

import "fmt"

func Solution(X int, A []int) int {
	// Implement your solution here
	if len(A) < X {
		return -1
	}

	var dist = make([]int, X)
	cnt := X
	for i, v := range A {
		if dist[v-1] == 0 {
			dist[v-1] = 1
			cnt--
			if cnt == 0 {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(Solution(5, []int{1, 3, 1, 4, 2, 3, 5, 4}))
}
