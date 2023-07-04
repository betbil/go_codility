package PermCheck

func Solution(A []int) int {
	// Implement your solution here
	length := len(A)
	cnt := length
	var aa = make([]int, length)
	for _, v := range A {
		if v > length || v < 1 || aa[v-1] > 0 {
			return 0
		}
		aa[v-1]++
		cnt--
	}
	if cnt == 0 {
		return 1
	}
	return 0
}
