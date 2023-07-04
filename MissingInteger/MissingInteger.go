package MissingInteger

// todobetul check this
func Solution(A []int) int {
	// Implement your solution here

	n := len(A)

	// Step 1: Mark the presence of positive integers
	present := make([]bool, n+1)
	for i := 0; i < n; i++ {
		if A[i] > 0 && A[i] <= n {
			present[A[i]] = true
		}
	}

	// Step 2: Find the smallest missing positive integer
	for i := 1; i <= n; i++ {
		if !present[i] {
			return i
		}
	}

	// If all positive integers are present, return n+1
	return n + 1
}
