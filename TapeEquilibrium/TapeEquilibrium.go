package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Solution(A []int) int {
	// Implement your solution here
	N := len(A)
	totalSum := 0

	// Calculate the total sum of the array
	for i := 0; i < N; i++ {
		totalSum += A[i]
	}

	//minDiff := totalSum
	minDiff := int(^uint(0) >> 1) // equivalent to math.MaxInt32
	fmt.Println("minDiff", minDiff, "^uint(0)", ^uint(0))

	leftSum := 0
	rightSum := totalSum

	// Iterate over the array and calculate the difference for each split point
	for i := 0; i < N-1; i++ {
		leftSum += A[i]
		rightSum -= A[i]

		diff := abs(leftSum - rightSum)

		fmt.Println("leftSum", leftSum, "rightSum", rightSum, "diff", diff)

		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

func main() {
	fmt.Println(Solution([]int{3, 1, 2, 4, 3}))
	fmt.Println(Solution([]int{1, 1}))
	fmt.Println(Solution([]int{-1000, 1000}))
}
