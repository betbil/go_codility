package main

import "fmt"

func moveZeroes(nums []int) {

	fmt.Println("nums:", nums)

	for i, j := 0, -1; i < len(nums) && j < len(nums); {
		if nums[i] != 0 {
			i++
		} else if j < len(nums) {
			if j == -1 {
				j = i + 1
				for j < len(nums) {
					fmt.Println("i", i, "j", j)
					if nums[j] == 0 {
						j++
					} else {
						break
					}
				}
			}
			fmt.Println("hey2")
			if j < len(nums) {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					i++
				}
				j++
				fmt.Println("hey")
				fmt.Println("i2", i, "j2", j)
			}

		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
