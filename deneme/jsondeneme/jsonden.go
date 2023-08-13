package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CaseInsensitivePerson struct {
	Name string
	Age  int
}

func (p *CaseInsensitivePerson) UnmarshalJSON(data []byte) error {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	for k, v := range m {
		if strings.EqualFold(k, "name") {
			if err := json.Unmarshal(v, &p.Name); err != nil {
				return err
			}
		} else if strings.EqualFold(k, "age") {
			if err := json.Unmarshal(v, &p.Age); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	data := `{"NAME": "Alice", "AGE": 30}`

	var p CaseInsensitivePerson
	err := json.Unmarshal([]byte(data), &p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
