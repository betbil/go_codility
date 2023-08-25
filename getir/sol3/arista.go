package main

import (
	"fmt"
	"strings"
)

func FormatStringpre(s string) string {
	cleanedString := ""
	for _, char := range s {
		if char != ' ' && char != '-' {
			cleanedString += string(char)
		}
	}

	formattedString := ""
	for i := 0; i < len(cleanedString); i += 3 {
		// If there are more than three characters left, append a group of three followed by a space
		if len(cleanedString)-i > 3 {
			formattedString += cleanedString[i:i+3] + " "
		} else {
			// If three or fewer characters are left, append the remaining characters without a trailing space
			formattedString += cleanedString[i:]
			break
		}
	}

	return strings.TrimSpace(formattedString)
}

func FormatString(s string) string {
	updatedFromSpaces := ""
	for _, char := range s {
		if char != ' ' && char != '-' {
			updatedFromSpaces += string(char)
		}
	}

	formattedSt := ""
	remaining := len(updatedFromSpaces)
	for i := 0; i < len(updatedFromSpaces); i += 3 {
		// to update last two
		if remaining == 2 || remaining == 4 {
			formattedSt += updatedFromSpaces[i:i+2] + " "
			i -= 1
			remaining -= 2
		} else {
			blockSize := 3
			/*if remaining == 3 || remaining == 5 {
				blockSize = 2
			}*/
			formattedSt += updatedFromSpaces[i:i+blockSize] + " "
			remaining -= blockSize
		}
	}

	return strings.TrimSpace(formattedSt)
}

func main() {

	fmt.Println(FormatString("00-44  48 5555 8361"))
	fmt.Println(FormatString("0 - 22 1985--324"))
	fmt.Println(FormatString("555372654"))
	fmt.Println(FormatString("00444855558361"))
}
