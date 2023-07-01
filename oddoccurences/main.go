package main

import (
	"fmt"
	"sort"
	"strconv"
)

func Solutionxor(A []int) int {
	// Implement your solution here
	if len(A)&1 == 0 {
		return -1
	}

	found := 0
	for _, v := range A {
		found ^= v
	}
	return found
}

func SolutionSort(A []int) int {
	// Implement your solution here
	if len(A)&1 == 0 {
		return -1
	}

	sort.Ints(A)
	low := 0
	high := len(A) - 1

	fmt.Println("A", A)
	for low < high {
		mid := (low + high) / 2
		fmt.Println("pre low", low, "high", high, "mid", mid)
		prenid := mid
		if mid&1 == 1 {
			mid--
		}

		fmt.Println(fmt.Sprintf("A[%d] = %d, A[%d] = %d", mid, A[mid], mid+1, A[mid+1]))
		if A[mid] == A[mid+1] {
			low = mid + 2
		} else {
			high = prenid
		}
		fmt.Println("after low", low, "high", high, "mid", mid)
	}
	return A[low]
}

func main() {
	a := 9
	b := 9
	c := 5
	fmt.Println("a tp binary", strconv.FormatInt(int64(a), 2))
	fmt.Println("b tp binary", strconv.FormatInt(int64(b), 2))
	fmt.Println(a ^ b ^ c)
	fmt.Println(Solutionxor([]int{9, 3, 9, 3, 9, 7, 9}))
	fmt.Println(SolutionSort([]int{9, 3, 9, 3, 9, 7, 9})) //3 3 7 9 9 9 9
	fmt.Println(SolutionSort([]int{9, 3, 9, 3, 7}))       //3 3 7 9 9 9 9
	hatali := []int{-8034843474017254924, -5777498374867451097, -2289411683948202627, -947120291822621557, 2049642480346318233, 2120283203010594819, 4103990856394083237, 4527516735047530362, 8987734566510693546}
	fmt.Println("xor", Solutionxor(hatali))
	fmt.Println(SolutionSort(hatali))
}
