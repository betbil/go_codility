package main

import (
	"fmt"
	"strconv"
)

func Solution(N int) int {
	// Implement your solution here
	num := N
	max := 0
	i := -1
	for num != 0 {
		currBit := num & 1
		if currBit == 1 {
			i = 0
		} else if i >= 0 {
			i++
			if i > max {
				max = i
			}
		}
		num = num >> 1
	}
	return max
}

func main() {
	num := 45
	binary := strconv.FormatInt(int64(num), 2)
	fmt.Println(binary)
	fmt.Println(Solution(num))
	num = 24
	binary = strconv.FormatInt(int64(num), 2)
	fmt.Println(binary)
	fmt.Println(Solution(num))
}
