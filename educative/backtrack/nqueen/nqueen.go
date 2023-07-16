package main

// todobetul check
func solveNQueens(n int) int {
	// Write your code here
	// Your code will replace this placeholder return statement

	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}

	count := 0

	placeQueen(0, &count, board)
	return count
}

func placeQueen(col int, count *int, board [][]bool) {

	if col == len(board) {
		*count++
	}

	for row := 0; row < len(board); row++ {
		if isSafe(row, col, board) {
			board[row][col] = true
			placeQueen(col+1, count, board)
			board[row][col] = false
		}
	}
}

func isSafe(row int, col int, board [][]bool) bool {
	for i := 0; i < col; i++ {
		if board[row][i] {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}

	for i, j := row, col; i < len(board) && j >= 0; i, j = i+1, j-1 {
		if board[i][j] {
			return false
		}
	}

	return true
}
