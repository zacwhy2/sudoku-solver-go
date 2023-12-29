package main

import (
	"fmt"
	"time"
)

func main() {
	s := `
		..7 .4. 28.
		..8 .12 9..
		.2. ... 1..

		... 76. ...
		16. ... ..5
		.8. ... ...

		.7. ... .6.
		... ..4 ..1
		..9 .81 3.7
	`

	onUpdate := func(board [][]byte) {
		time.Sleep(1 * time.Millisecond)
		printBoard(board)
	}
	Solve(StringToBoard(s), onUpdate)
}

// Prints the current board
func printBoard(board [][]byte) {
	clearScreen()
	for i, s := range board {
		for j, c := range s {
			fmt.Printf("%c", c)
			if (j+1)%3 == 0 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
}

func clearScreen() {
	fmt.Println("\033[2J\033[H") // clear screen
}
