package helper

import (
	"fmt"
)

// todobetul güzel
func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0

	for num := range numSet {
		if !numSet[num-1] { // Check if it could be the start of a sequence
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] { // Try to extend the sequence
				currentNum++
				currentStreak++
			}

			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums)) // Output: 4
}
