package main

import "fmt"

func firstUniqChar(s string) int {
	var starray [26]int
	fmt.Print("s, len(s): ", s, " ", len(s), "\n")
	fmt.Println("starray: ", starray)
	for _, v := range s {
		starray[v-'a']++
	}
	for i, v := range s {
		if starray[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}
