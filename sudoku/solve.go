package sudoku

import (
	"math"
)

// Gets the goods
func (b *Board) Solve() bool {
	// TODO: Check first if board is solved and if so, exit
	for y := 0; y < b.Size; y++ {
		for x := 0; x < b.Size; x++ {
			curr := &b.Cells[b.getIndex(x, y)]
			if curr.Val == 0 {
				return b.solveCell(curr)
			}
		}
	}
	return true
}

func (b *Board) solveCell(c *Cell) bool {
	for i := 1; i <= b.Size; i++ {
		if b.checkAll(*c, i) {
			c.Val = i
			b.Steps++
			if b.Solve() {
				return true
			}
			c.Val = 0
		}
	}
	return false
}

// An alternative take on the same fundemental recursive backtracking idea
// func (b *Board) Solve() bool {
// 	curr := b.next()
// 	if curr == nil {
// 		return true
// 	}

// 	for i := 1; i < 10; i++ {
// 		if b.checkAll(*curr, i) {
// 			curr.Val = i
// 			if b.Solve() {
// 				return true
// 			}
// 			curr.Val = 0
// 		}
// 	}
// 	return false
// }

// // Returns the next zero value in the board
// func (b *Board) next() *Cell {
// 	for i := range b.Cells {
// 		if b.Cells[i].Val == 0 {
// 			return &b.Cells[i]
// 		}
// 	}
// 	return nil
// }

// Checks a value against a cell's residing row, column, and block
func (b *Board) checkAll(c Cell, val int) bool {
	rowOk := check(b.getRow(c), val)
	colOk := check(b.getCol(c), val)
	blockOk := check(b.getBlock(c), val)
	return (rowOk && colOk && blockOk)
}

// Checks if a given value is in a slice of cells
func check(block []Cell, val int) bool {
	for i := range block {
		if block[i].Val == val {
			return false
		}
	}
	return true
}

// TODO: get hacky bits working on getRow and getCol to improve speed for larger puzzles
// Returns the row a given cell resides in
func (b *Board) getRow(c Cell) []Cell {
	result := []Cell{}
	for i := range b.Cells {
		curr := b.Cells[i]
		if curr.Y == c.Y {
			result = append(result, curr)
		}
	}
	return result
	// index := (c.Y - 1) * b.Size
	// result := b.Cells[index:(index + b.Size)]
	// return result
}

// Returns the column a given cell resides in
func (b *Board) getCol(c Cell) []Cell {
	result := []Cell{}
	for i := range b.Cells {
		curr := b.Cells[i]
		if curr.X == c.X {
			result = append(result, curr)
		}
	}
	return result

	// index := c.X + 1
	// result := []Cell{b.Cells[index]}
	// for i := 0; i < (b.Size - 1); i++ {
	// 	index += b.Size
	// 	result = append(result, b.Cells[index])
	// }
	// return result
}

// Returns the block a given cell resides in
func (b *Board) getBlock(c Cell) []Cell {
	result := []Cell{}
	mult := int(math.Sqrt(float64(b.Size)))
	xBlock := (c.X / mult) * mult
	yBlock := (c.Y / mult) * mult

	for x := (xBlock); x < (xBlock + mult); x++ {
		for y := (yBlock); y < (yBlock + mult); y++ {
			result = append(result, b.Cells[b.getIndex(x, y)])
		}
	}
	return result
}
