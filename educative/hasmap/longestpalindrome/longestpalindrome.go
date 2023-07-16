package main

func longestPalindrome(palString string) int {
	// Your code will replace the placeholder return statement.
	if len(palString) == 0 {
		return 0
	}
	occurences := make(map[rune]int)
	oddnum := 0
	evenNum := 0

	for _, v := range palString {
		occurences[v]++
		if occurences[v]%2 == 0 {
			evenNum++
			oddnum--
		} else {
			oddnum++
		}
	}
	if oddnum > 0 {
		return evenNum*2 + 1
	}
	return evenNum * 2
}
