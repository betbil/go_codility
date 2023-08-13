package main

import (
	"fmt"
)

func main() {
	s := "Hello, 世界"

	printBytesAndRunes(s)
	printBytesAndRunes("den123")
}

func printBytesAndRunes(s string) {
	fmt.Println(s)
	fmt.Println("len(s)", len(s))
	for _, r := range s {
		fmt.Printf("%c ", r) // Prints each character, including multi-byte characters
	}
	fmt.Println("Bytes:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d. %x ", i, s[i]) // Prints each byte in hex
	}
	fmt.Println("\nRunes:")
	for i, r := range s {
		fmt.Printf("%d. %x ", i, r) // Prints each rune (Unicode code point) in hex
	}
	rs := []rune(s)
	fmt.Println("\nRunes:", rs, "len(rs)", len(rs))
	for i, r := range rs {
		fmt.Printf("%d. %x ", i, r) // Prints each rune (Unicode code point) in hex
	}
	rs2 := []rune("Hello, 世界")
	fmt.Println("\nRunes:", rs2, "len(rs)", len(rs2))
	bytes := []byte(s)
	bytes2 := []byte("Hello, 世界")
	fmt.Println("\nbytes:", bytes, "len(rs)", len(bytes))
	fmt.Println("\nbytes2:", bytes2, "len(rs)", len(bytes2))

}
