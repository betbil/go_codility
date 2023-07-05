package main

import "fmt"

// todobetul	check
func Solution(A int, B int, K int) int {
	// Implement your solution here
	fmt.Println("A", A, "B", B, "K", K)
	if A == B {
		if A%K == 0 {
			return 1
		}
		return 0
	}
	cnt := B / K
	if A > 0 {
		cnt -= (A - 1) / K
	} else {
		cnt++
	}
	return cnt
}

func main() {
	fmt.Println(Solution(6, 11, 2))
	fmt.Println(Solution(0, 0, 11))
	fmt.Println(Solution(11, 345, 17))
	fmt.Println(Solution(10, 10, 5))
	fmt.Println(Solution(0, 0, 1))
}
