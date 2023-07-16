package main

import "fmt"

// todobetul check
func singleNonDuplicate(nums []int) int {

	// replace this placeholder return statement with your code
	if len(nums) == 0 {
		return -1
	}
	low := 0
	high := len(nums) - 1
	fmt.Println("low", low, "high", high)

	for low < high {

		mid := low + (high-low)/2
		if mid%2 == 1 {
			mid--
		}

		if nums[mid] != nums[mid+1] {
			high = mid
		} else {
			low = mid + 2
		}
		fmt.Println("low", low, "high", high, "mid", mid)
	}

	return nums[low]
}

func main() {
	fmt.Println("singleNonDuplicate", singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println("singleNonDuplicate", singleNonDuplicate([]int{1, 1, 2, 3, 3}))
}
