package passingcars

func Solution(A []int) int {
	// Implement your solution here
	total := 0
	zeros := 0
	for _, v := range A {
		if v == 0 {
			zeros++
		} else {
			total += zeros
		}
		if total > 1000000000 {
			return -1
		}
	}
	return total
}
