package main

import "fmt"

func permuteWord(word string) []string {
	data := []rune(word)
	var result []string
	permuteHelper(data, 0, len(data), &result)
	return result
}

func permuteHelper(data []rune, start int, end int, result *[]string) {
	fmt.Println("inn permuteHelper", string(data), start, end, result)
	if start == end-1 {
		*result = append(*result, string(data))
		return
	}

	for i := start; i < end; i++ {
		data[start], data[i] = data[i], data[start]
		//fmt.Println("before call permuteHelper", string(data), start, end, result)
		permuteHelper(data, start+1, end, result)
		data[start], data[i] = data[i], data[start]
		//fmt.Println("after call permuteHelper", string(data), start, end, result)
	}
}

func main() {
	fmt.Println("permuteword", permuteWord("ABC"))
	//fmt.Println("permuteword", permuteWord("ABCD"))
}
