package helper

import (
	"fmt"
	"strconv"
)

func Solution(N int) int {
	// Implement your solution here
	//16 8 4 2 1
	i := 0
	for {
		isOdd := (N & 1) == 1
		if isOdd {
			break
		}
		N = N >> 1
		i++
	}
	return i
}

func main() {
	num := 45
	binary := strconv.FormatInt(int64(num), 2)
	fmt.Println(binary)
	fmt.Println(Solution(num))
	num = 24
	binary = strconv.FormatInt(int64(num), 2)
	fmt.Println(binary)
	fmt.Println(Solution(num))
}
