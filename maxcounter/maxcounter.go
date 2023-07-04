package maxcounter

// todobetul check this
func Solution(N int, A []int) []int {
	// Implement your solution here
	var result = make([]int, N)
	max := 0
	maxi := 0
	for _, v := range A {
		if v <= N {
			if maxi > result[v-1] {
				result[v-1] = maxi
			}

			result[v-1]++
			val := result[v-1]
			if val > max {
				max = val
			}
		} else {
			maxi = max
		}
	}

	for i, _ := range result {
		if maxi > result[i] {
			result[i] = maxi
		}
	}

	return result
}
