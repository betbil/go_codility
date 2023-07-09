package main

import (
	"fmt"
	"strings"
)

// todobetul check
func reverseWords(sentence string) string {
	// Your code will replace the placeholder return statement
	fmt.Println("received sentence", sentence)
	st2 := strings.ReplaceAll(string(sentence), " ", "")
	fmt.Println("st2", st2)
	length := len(sentence)
	l := 0
	h := length - 1
	bytes := []byte(sentence)

	for l < h {
		bytes[l], bytes[h] = sentence[h], sentence[l]
		l++
		h--
	}
	fmt.Println("before reversed sentence", string(bytes))

	l = 0
	h = 0
	for i, b := range bytes {
		fmt.Println("i, b ", i, string(b))
		if b == ' ' {
			h = i - 1
			fmt.Println("l, h ", l, h)
			for l < h {
				bytes[l], bytes[h] = bytes[h], bytes[l]
				l++
				h--
				fmt.Println("updated bytes:", string(bytes))
			}
			l = i + 1
		} else if i == length-1 {
			h = i
			fmt.Println("last l, h ", l, h)
			for l < h {
				bytes[l], bytes[h] = bytes[h], bytes[l]
				l++
				h--
			}
		}
	}

	fmt.Println("reversed sentence", string(bytes))

	/*

		substring := strings.Split(sentence, " ")
		substrings2 := strings.Fields(sentence)
		fmt.Println("substring", substring)
		fmt.Println("substring2", substrings2)

		listst := list.New()
		for _, s := range substring {
			listst.PushFront(s)
		}
		println(listst)


	*/
	return string(bytes)
}

func main() {
	fmt.Println(reverseWords("Hello World deneme 1234  5678"))
}
