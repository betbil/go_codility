package main

import (
	"fmt"
	"sort"
)

func findSumOfThree(nums []int, target int) bool {
	// Replace this placeholder return statement with your code
	length := len(nums)
	if length < 3 {
		return false
	}

	sort.Ints(nums)
	for i := 0; i < length; i++ {
		l := i + 1
		h := length - 1
		for l < h {
			if nums[i]+nums[l]+nums[h] == target {
				fmt.Println(nums[i], nums[l], nums[h])
				return true
			} else if nums[i]+nums[l]+nums[h] < target {
				l++
			} else {
				h--
			}
		}
	}

	return false
}

func main() {
	fmt.Println(findSumOfThree([]int{1, 2, 6, 4, 5}, 9))
}
