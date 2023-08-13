package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	result := make([][]int, 0)
	for _, n1 := range nums1 {
		m1[n1] = true
	}

	for _, n2 := range nums2 {
		_, ok := m1[n2]
		if ok {
			m1[n2] = false
		} else {
			m2[n2] = true
		}
	}

	m1r := make([]int, 0)
	for k, v := range m1 {
		if v {
			m1r = append(m1r, k)
		}
	}
	result = append(result, m1r)
	m2r := make([]int, 0)
	for k, v := range m2 {
		if v {
			m2r = append(m2r, k)
		}
	}

	result = append(result, m2r)

	return result

}

func main() {
	nums1 := []int{5, 2, 3}
	nums2 := []int{2, 4, 6}
	findDifference(nums1, nums2)
}
