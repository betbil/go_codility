package main

import "fmt"

func Solution(N int) int {
	if N == 0 {
		return 0
	}

	currBit := 0
	i := 0
	tmp := 0
	num := N
	num1 := 0
	num2 := 0
	prn1 := 0
	prn2 := 0

	for num != 0 {
		currBit = num & 1
		if currBit == 1 {
			tmp = 1 << i
			if prn1 == 0 {
				num1 = num1 | tmp
				prn1 = 1
				prn2 = 0
			} else if prn2 == 0 {
				num2 = num2 | tmp
				prn2 = 1
				prn1 = 0
			} else {
				return -1
			}
		}
		i = i + 1
		num = num >> 1
	}

	fmt.Println("num1", num1, "num2", num2)

	if num1 > 0 {
		return num1
	}
	if num2 > 0 {
		return num2
	}

	return -1

}

func main() {
	fmt.Println(Solution(1031))
}
