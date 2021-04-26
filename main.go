package main

import (
	// "fmt"
	"sudoku-solver/sudoku"
)

func main() {
	board := sudoku.Create(9, 1)
	board.Display()

	// cell := board.Cells[21]
	// fmt.Printf("cell: %v\n", cell.Val)
	// row := board.GetRow(cell)
	// fmt.Printf("row: ")
	// Print(row)
	// col := board.GetCol(cell)
	// fmt.Printf("col: ")
	// Print(col)
	// block := board.GetBlock(cell)
	// fmt.Printf("block: ")
	// Print(block)

	board.Solve()
	board.Display()
}

// func Print(c []sudoku.Cell) {
// 	for i := range c {
// 		fmt.Printf("%v, ", c[i].Val)
// 	}
// 	fmt.Printf("\n")
// }
