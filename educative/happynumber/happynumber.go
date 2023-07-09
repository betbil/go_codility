package main

import (
	"fmt"
)

// todobetul check
func isHappy(num int) bool {
	// write your code here
	// your code will replace this placeholder return statement
	fmt.Println("isHappy num", num)
	if num == 1 {
		return true
	}
	slow := num
	fast := sumofDigits(num)

	for fast != 1 {
		if slow == fast {
			return false
		}
		slow = sumofDigits(slow)
		fast = sumofDigits(sumofDigits(fast))
	}

	return true
}

func sumofDigits(num int) int {
	sum := 0
	for num != 0 {
		sum += (num % 10) * (num % 10)
		num /= 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy(23))
	fmt.Println(isHappy(4))
	fmt.Println(isHappy(1))
}
