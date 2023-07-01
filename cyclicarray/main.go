package main

import "fmt"

func Solution(A []int, K int) []int {
	// Implement your solution here
	length := len(A)
	if length == 0 {
		return A
	}
	B := make([]int, length)
	for i, v := range A {
		B[(i+K)%length] = v
	}

	return B
}

func main() {
	fmt.Println(Solution([]int{3, 8, 9, 7, 6}, 3))
}
