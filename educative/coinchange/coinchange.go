package main

import "fmt"

func chatcanPartition(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}

	// If the total sum is odd, it cannot be partitioned into two equal subsets
	if total%2 != 0 {
		return false
	}

	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
			fmt.Println(fmt.Sprintf("dp[%d](%v) = dp[%d](%v) || dp[%d-%d](%v)", j, dp[j], j, dp[j], j, nums[i], dp[j-nums[i]]), "\n")
		}
	}

	return dp[target]
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(chatcanPartition(nums)) // true

	nums2 := []int{1, 2, 3, 5}
	fmt.Println(chatcanPartition(nums2)) // false
}
