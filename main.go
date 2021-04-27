package main

import (
	"sudoku-solver/sudoku"
)

func main() {
	board := sudoku.Create(9, 3)
	board.Display()
	board.Solve()
	board.Display()
}
