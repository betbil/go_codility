package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	s := "Hello, 世界"
	// s := "aliveli"
	fmt.Println("s", s)
	runes := []rune(s)
	//runes := []byte(s) //bhyte aliveli de çalışır ama hello da çalışmaz
	l := 0
	h := len(runes) - 1
	for l < h {
		runes[l], runes[h] = runes[h], runes[l]
		l++
		h--
	}
	s = string(runes)
	fmt.Println("s", s)

}
