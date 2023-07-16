package main

import (
	"fmt"
)

// todobetul check
func chatgbtcanJump(nums []int) bool {
	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

func edujumpGame(nums []int) bool {
	targetNumIndex := len(nums) - 1
	for i := len(nums) - 2; i > -1; i-- {
		if targetNumIndex <= i+nums[i] {
			targetNumIndex = i
		}
	}
	if targetNumIndex == 0 {
		return true
	}
	return false
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(chatgbtcanJump(nums)) // Expected output: true

	nums = []int{3, 2, 1, 0, 4}
	var nums2 []int
	fmt.Println(chatgbtcanJump(nums))  // Expected output: false
	fmt.Println(chatgbtcanJump(nums2)) // Expected output: false
}
