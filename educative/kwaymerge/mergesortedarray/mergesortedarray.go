package main

import "fmt"

// todobetul check
func mergeSorted(nums1 []int, m int, nums2 []int, n int) []int {
	// Write your code here
	// Your code will replace this placeholder return statement
	//len(num1) m+n
	fmt.Println("nums1", nums1, "m", m)
	fmt.Println("nums2", nums2, "n", n)
	if len(nums1) < m+n || len(nums2) < n {
		return nil
	}

	p1 := m - 1
	p2 := n - 1
	p3 := m + n - 1

	for p3 >= 0 {
		if p1 >= 0 && p2 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p3] = nums1[p1]
			p1--
		} else if p2 >= 0 {
			nums1[p3] = nums2[p2]
			p2--
		}
		p3--
	}

	return nums1
}

func main() {
	fmt.Println("serted", mergeSorted([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
	fmt.Println("sorted", mergeSorted([]int{6, 7, 8, 9, 10, 0, 0, 0, 0, 0}, 5, []int{1, 2, 3, 4, 5}, 5))

}
