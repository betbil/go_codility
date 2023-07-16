package main

import "fmt"

// todobetul check
func binarySearch(nums []int, target int) int {
	// Write your code here
	fmt.Println("received nums", nums, "target", target)
	l := 0
	h := len(nums) - 1
	if h < 0 {
		return -1
	}
	for l < h {
		mid := (l + h) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func edubinarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		// Finding the mid using integer division
		//todobetul unutma ve low <= high olduğu için herşery bu forda olmuş oluyor dışarıya -1 döndürüyor
		mid := low + (high-low)/2

		// Target value is present at the middle of the slice
		if nums[mid] == target {
			return mid
		}
		// Target value is present in the low subarray
		if target < nums[mid] {
			high = mid - 1
		}
		// Target value is present in the high subarray
		if target > nums[mid] {
			low = mid + 1
		}
	}

	// Target value is not present in the slice
	return -1
}

func main() {
	fmt.Println("binarySearch", binarySearch([]int{1, 2, 3, 4, 5, 6, 7}, 5))
	fmt.Println("binarySearch", binarySearch([]int{1, 2, 3, 4, 5, 6, 8}, 7))
	fmt.Println("binarySearch", binarySearch([]int{}, 3))
}
