package main

import (
	"sudoku-solver/sudoku"
)

func main() {
	board := sudoku.CreateBoard(4, 1)
	body := board.Fetch()
	res := new(sudoku.Response)

	sudoku.Must(res.FromJson(body))
	board.Cells = res.Squares

	board.Print()
}
