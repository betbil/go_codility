package main

import (
	"fmt"
	"sort"
)

// todobetul check
func sortColors(colors []int) []int {

	// Write your code here
	length := len(colors)
	if length == 0 {
		return colors
	}

	red, white, blue := 0, 0, length-1 //0,1,2

	for white <= blue {
		if colors[white] == 0 {
			colors[red], colors[white] = colors[white], colors[red]
			red++
			white++
		} else if colors[white] == 1 {
			white++
		} else {
			colors[white], colors[blue] = colors[blue], colors[white]
			blue--
		}
	}

	return colors
}

func main() {
	fmt.Println(sortColors([]int{0, 1, 2, 0, 1, 2}))
	input := []int{0, 1, 2, 0, 1, 2}
	sort.Ints(input)
	fmt.Println(input)
}
