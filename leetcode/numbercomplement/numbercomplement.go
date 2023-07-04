package main

import (
	"fmt"
)

func findComplement(num int) int {
	allone := 1
	for allone < num {
		allone = (allone << 1) | 1
	}
	return allone ^ num
}

func main() {
	fmt.Printf("5 cmpl: %b\n", findComplement(5))
	fmt.Printf("1 cmpl: %b\n", findComplement(1))
	fmt.Printf("5 cmpl: %d\n", findComplement(5))
}
