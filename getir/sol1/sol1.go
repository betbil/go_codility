package main

import (
	"fmt"
	"math/big"
	"time"
)

func Solution(A int, B int) int {
	// Implement your solution here
	return solutionWithInt64(int64(A), int64(B))

}

func Solution2(A int, B int) int {
	// Implement your solution here
	return solutionWithInt64(int64(A), int64(B))

}

func solutionWithInt64(A int64, B int64) int {
	mult := big.NewInt(0).Mul(big.NewInt(A), big.NewInt(B))

	// Convert to binary string
	binaryStr := fmt.Sprintf("%b", mult)

	count := 0
	for _, ch := range binaryStr {
		if ch == '1' {
			count++
		}
	}

	time.Sleep(2 * time.Second)
	return count
}

func main() {
	fmt.Println(Solution(0, 7))
	fmt.Println(Solution(2, 5))
	fmt.Println(Solution(100000000, 100000000))
}
