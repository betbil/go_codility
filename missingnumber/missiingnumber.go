package main

import "fmt"

func SolutionMissing(A []int) int {
	// Implement your solution here
	fmt.Println("arraylenghth", len(A))
	n := len(A) + 1
	if n == 1 {
		return 1
	}
	sum := n * (n + 1) / 2
	for _, v := range A {
		sum -= v
	}
	return sum
}

func main() {
	fmt.Println(SolutionMissing([]int{2, 3, 1, 5}))
	fmt.Println(SolutionMissing([]int{}))
}
