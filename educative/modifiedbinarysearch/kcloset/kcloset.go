package main

import "fmt"

func findClosestElementsNotworking(nums []int, k int, target int) []int {

	// your code will replace this placeholder return statement
	if len(nums) == 0 {
		return nums
	}

	if len(nums) < k {
		k = len(nums)
	}

	if nums[0] >= target {
		return nums[:k]
	}
	if nums[len(nums)-1] <= target {
		return nums[len(nums)-k:]
	}

	// find the closest element to target
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			start = mid
			break
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	// start is the closest element to target
	// find the k closest elements to start
	left, right := start, start+1
	for right-left-1 < k && (left >= 0 || right < len(nums)) {
		if left < 0 {
			right++
			continue
		}
		if right > len(nums)-1 {
			left--
			continue
		}
		if target-nums[left] <= nums[right]-target {
			left--
		} else {
			right++
		}
	}
	return make([]int, 0)
}

// todobetul güzel çözüm
func chatgbtfindClosestElements(nums []int, target, k int) []int {
	left := 0
	right := len(nums) - 1

	// Binary search to find the closest element to the target
	for right-left+1 > k {
		if target-nums[left] <= nums[right]-target {
			right--
		} else {
			left++
		}
	}

	// Return the k closest elements
	return nums[left : right+1]
}

func main() {
	nums := []int{-5, 1, 2}
	var nums2 []int
	target := 0
	k := 1

	fmt.Println(chatgbtfindClosestElements(nums, target, k))  // Output: [2 3 4]
	fmt.Println(chatgbtfindClosestElements(nums2, target, k)) // Output: [2 3 4]
}
