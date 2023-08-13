package main

func numIslands(grid [][]byte) int {

	// Your code will replace this placeholder return statement
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if dfs(i, j, grid) {
				count++
			}
		}
	}
	return -1
}

func checkValidty(i, j int, grid [][]byte) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] != '0'
}

func dfs(i, j int, grid [][]byte) bool {
	if !checkValidty(i, j, grid) {
		return false
	}
	grid[i][j] = '0'
	dfs(i+1, j, grid)
	dfs(i-1, j, grid)
	dfs(i, j+1, grid)
	dfs(i, j-1, grid)
	return true
}
