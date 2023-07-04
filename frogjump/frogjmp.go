package main

import "fmt"

func Solution(X int, Y int, D int) int {
	// Implement your solution here
	var cnt2 float64
	var cnt3 float64
	cnt := (Y - X) / D
	cnt2 = float64((Y - X) / D)
	cnt3 = float64(Y-X) / float64(D)
	fmt.Println("cnt", cnt)
	fmt.Println("cnt2", cnt2)
	fmt.Println("cnt3", cnt3)
	if (Y-X)%D != 0 {
		cnt++
	}
	fmt.Println("cnt", cnt)
	return cnt
}

func main() {
	fmt.Println(Solution(10, 85, 30))
	fmt.Println(Solution(5, 1000000, 2))
	fmt.Println(Solution(5, 5, 2))
}

func Solution2(X int, Y int, D int) int {
	// Implement your solution here
	CNT := 0
	for X < Y {
		X += D
		CNT++
	}
	return CNT
}
