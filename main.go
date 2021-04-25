package main

import (
	"sudoku-solver/sudoku"
)

func main() {
	board := sudoku.Create(4, 1)
	board.Print()
}
