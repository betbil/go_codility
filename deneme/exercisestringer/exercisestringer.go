package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String2() string {
	bytes := make([]byte, 0, 7)
	for i := 0; i < 3; i++ {
		bytes = append(bytes, ip[i])
		bytes = append(bytes, '.')
	}
	bytes = append(bytes, ip[3])
	s := string(bytes)
	return s
}
func (ip IPAddr) String() string {
	bytes := make([]rune, 0, 7)
	for i := 0; i < 3; i++ {
		bytes = append(bytes, []rune(strconv.Itoa(int(ip[i])))...)
		bytes = append(bytes, '.')
	}
	bytes = append(bytes, []rune(strconv.Itoa(int(ip[3])))...)
	return string(bytes)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	ss := "127.0.0.1"
	runes := []rune(ss)
	for i, v := range ss {
		fmt.Println(i, v)
	}

	for _, v := range runes {
		fmt.Println(string(v))
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
