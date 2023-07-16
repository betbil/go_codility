package main

import "fmt"

// Tip: You may use some of the code templates provided
// in the support files

// todobetul check why backtracking important
func wordSearch(grid [][]string, word string) bool {
	// Write your code here
	// Your code will replace this placeholder return statement
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if wordsearchrecur(grid, visited, word, 0, r, c) {
				return true
			}
		}
	}

	return false
}

func wordsearchrecur(grid [][]string, visited [][]bool, word string, wi, r, c int) bool {
	if wi == len(word) {
		return true
	}

	if !checkIndexValidity(grid, r, c) {
		return false
	}

	if visited[r][c] {
		return false
	}
	//fmt.Printf("word[%d]:%d, grid[%d][%d][0]:%d \n", wi, word[wi], r, c, grid[r][c][0])
	if word[wi] != grid[r][c][0] {
		return false
	}

	visited[r][c] = true
	result := wordsearchrecur(grid, visited, word, wi+1, r+1, c) ||
		wordsearchrecur(grid, visited, word, wi+1, r-1, c) ||
		wordsearchrecur(grid, visited, word, wi+1, r, c+1) ||
		wordsearchrecur(grid, visited, word, wi+1, r, c-1)
	visited[r][c] = false
	return result
}

func checkIndexValidity(grid [][]string, r, c int) bool {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
		return false
	}
	return true
}

func main() {
	//fmt.Println(wordSearch([][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}, "ABCCED"))
	fmt.Println(wordSearch([][]string{{"N", "W", "L", "I", "M"}, {"V", "I", "L", "Q", "O"}, {"O", "L", "A", "T", "O"}, {"R", "T", "A", "I", "N"}, {"O", "I", "T", "N", "C"}}, "LATIN"))
	fmt.Println(wordSearch([][]string{{"a", "b", "a"}, {"c", "a", "a"}, {"c", "c", "a"}}, "abaca"))
	fmt.Println(wordSearch([][]string{{"a", "x", "x"}, {"c", "b", "b"}, {"c", "c", "a"}}, "xbbc"))
	/*
		A x x
		C b b
		C C A

	*/
	/*
		{"N", "W", "L", "I", "M"}
		{"V", "I", "L", "Q", "O"}
		{"O", "L", "A", "T", "O"}
		{"R", "T", "A", "I", "N"}
		{"O", "I", "T", "N", "C"}}

	*/
}
