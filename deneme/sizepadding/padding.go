package main

import (
	"fmt"
	"unsafe"
)

type StructA struct {
	a bool
	b int64
	c bool
}

type StructB struct {
	a bool
	c bool
	e float32
	d int
	b int64
	f float64
}

func main() {
	sb := StructB{}
	fmt.Println("Size of sb.d int:", unsafe.Sizeof(sb.d))
	fmt.Println("Size of sb.b int64:", unsafe.Sizeof(sb.b))
	fmt.Println("Size of sb.a bool:", unsafe.Sizeof(sb.a))
	fmt.Println("Size of sb.e float32:", unsafe.Sizeof(sb.e))
	fmt.Println("Size of sb.f float64:", unsafe.Sizeof(sb.f))
	fmt.Println("Size of StructA:", unsafe.Sizeof(StructA{}))
	fmt.Println("Size of StructB:", unsafe.Sizeof(StructB{}))
}
