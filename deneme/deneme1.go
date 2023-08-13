package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	jsonData := []byte(`{"Name":"Alice","age":30}`)

	var p Person
	err := json.Unmarshal(jsonData, &p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p)
}
