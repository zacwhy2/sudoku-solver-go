package main

import (
	"strings"
)

// Solve returns the solution given an initial board
func Solve(board [][]byte, onUpdate func([][]byte)) [][]byte {
	solve(board, 0, 0, onUpdate)
	return board
}

func solve(board [][]byte, r int, c int, onUpdate func([][]byte)) bool {
	if r == 9 { // last row
		return true
	}

	if c == 9 { // last column
		return solve(board, r+1, 0, onUpdate)
	}

	if board[r][c] != '.' {
		return solve(board, r, c+1, onUpdate)
	}

	for ch := '1'; ch <= '9'; ch++ {
		if !isValid(board, r, c, byte(ch)) {
			continue
		}

		board[r][c] = byte(ch)
		onUpdate(board)

		if solve(board, r, c+1, onUpdate) {
			return true
		}

		board[r][c] = '.'
		onUpdate(board)
	}

	return false
}

func isValid(board [][]byte, r int, c int, n byte) bool {
	for i := 0; i < 9; i++ {
		// check row
		if board[r][i] == n {
			return false
		}
		// check column
		if board[i][c] == n {
			return false
		}
		// check 3*3 square
		if board[(r/3)*3+i/3][(c/3)*3+i%3] == n {
			return false
		}
	}
	return true
}

// StringToBoard converts string s to 9x9 number board
func StringToBoard(s string) [][]byte {
	board := make([][]byte, 0)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\t", "")
		line = strings.ReplaceAll(line, " ", "")
		if line == "" {
			continue
		}
		board = append(board, []byte(line))
	}
	return board
}
