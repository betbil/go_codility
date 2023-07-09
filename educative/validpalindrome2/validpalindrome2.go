// Valid Palindrome II
package main

import "fmt"

func validPalindrome(s string) bool {
	// Write your code here
	// Tip: You may use the code template provided
	// in the two_pointers.py file
	// your code will replace this placeholder return statement

	// 1. check if s is a palindrome
	// 2. if not, check if s[1:len(s)] or s[0:len(s)-1] is a palindrome
	// 3. if not, return false
	// 4. if yes, return true
	fmt.Println("validPalindrome", s)
	l := 0
	h := len(s) - 1
	fmt.Println("l", l, "h", h)
	for l < h {
		if s[l] != s[h] {
			fmt.Println("s", s, "not equal l", l, "h", h)
			return checkPalindrome(s[l+1:h+1]) || checkPalindrome(s[l:h])
		}
		l++
		h--
	}
	return true

}

func checkPalindrome(s string) bool {
	l := 0
	h := len(s) - 1

	fmt.Printf("s:%s h:%d \n", s, h)
	for l < h {
		if s[l] != s[h] {
			return false
		}
		l++
		h--
	}
	return true

}

func main() {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("abc"))
	fmt.Println(validPalindrome("ab"))
	fmt.Println(validPalindrome("a"))
	fmt.Println(validPalindrome("abccba"))
	fmt.Println(validPalindrome("abccbae"))

}
