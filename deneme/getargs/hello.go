package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("received args:", os.Args)
	if len(os.Args) < 2 {
		println("usage: %s name", os.Args[0])
		return
	}
	name := os.Args[1]
	fmt.Println("hello", name)
}
