package main

import (
	"fmt"
)

func findRepeatedSequences(s string, k int) []string {

	// Your code will replace this placeholder return statement
	stmap := make(map[string]int)
	var repeated []string
	for i := 0; i < len(s)-k+1; i++ {
		st := s[i : i+k]
		if _, ok := stmap[st]; ok {
			repeated = append(repeated, st)
		}
		stmap[st]++
	}
	return repeated
}

func main() {
	fmt.Println(findRepeatedSequences("awaglknagawunagwkwagl", 4))
	fmt.Println(findRepeatedSequences("aa1aa2aa", 2))
	/*
		input := "Hello, world!"

		// Create a new SHA256 hasher
		hasher := sha256.New()

		// Write the string to the hasher
		hasher.Write([]byte(input))

		// Get the hash value as a byte slice
		hashBytes := hasher.Sum(nil)

		// Convert the byte slice to a hexadecimal string
		hashString := fmt.Sprintf("%x", hashBytes)

		fmt.Println("hashBytes:", hashBytes)
		fmt.Println("hashString", hashString)

	*/
}
