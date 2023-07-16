package main

import "fmt"

func floodFill(grid [][]int, sr, sc, target int) [][]int {
	// Replace this placeholder return statement with your code
	fmt.Println(len(grid), len(grid[0]))
	if len(grid) == 0 {
		return grid
	}

	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}

	srcval := grid[sr][sc]

	floodFillrecur(grid, visited, sr, sc, srcval, target)

	return grid
}

func isValid(grid [][]int, visited [][]bool, r, c, src int) bool {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] != src || visited[r][c] {
		return false
	}
	return true
}
func floodFillrecur(grid [][]int, visited [][]bool, r, c, src, target int) {
	if !isValid(grid, visited, r, c, src) {
		return
	}

	grid[r][c] = target
	visited[r][c] = true

	floodFillrecur(grid, visited, r+1, c, src, target)
	floodFillrecur(grid, visited, r-1, c, src, target)
	floodFillrecur(grid, visited, r, c+1, src, target)
	floodFillrecur(grid, visited, r, c-1, src, target)

	//visited[r][c] = false
}

func main() {
	fmt.Println("floodFill", floodFill([][]int{{1, 2, 0, 4, 5}, {3, 1, 3, 6}, {7, 2, 1, 5}}, 1, 1, 4))
	//fmt.Println(len(grid), len(grid[0])) //3,5 len row sayısı todobetul check
}
