package sudoku

import (
	"math"
)

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
			if b.Solve() {
				return b.Solve()
			}
			c.Val = 0
		}
	}
	return false
}

func (b *Board) checkAll(c Cell, val int) bool {
	rowOk := b.checkRow(c, val)
	colOk := b.checkCol(c, val)
	blockOk := b.checkBlock(c, val)
	return (rowOk && colOk && blockOk)
}

func (b *Board) checkRow(c Cell, val int) bool {
	block := b.getRow(c)
	return check(block, val)
}

func (b *Board) checkCol(c Cell, val int) bool {
	block := b.getCol(c)
	return check(block, val)
}

func (b *Board) checkBlock(c Cell, val int) bool {
	block := b.getBlock(c)
	return check(block, val)
}

func check(block []Cell, val int) bool {
	for i := range block {
		if block[i].Val == val {
			return false
		}
	}
	return true
}

func (b *Board) getRow(c Cell) []Cell {
	start := c.Y * b.Size
	result := b.Cells[start:(start + b.Size)]
	return result
}

func (b *Board) getCol(c Cell) []Cell {
	result := []Cell{b.Cells[c.X]}
	for i := 0; i < b.Size; i++ {
		result = append(result, b.Cells[c.X+b.Size])
	}
	return result
}

// Returns the block a given cell resides in
func (b *Board) getBlock(c Cell) []Cell {
	result := []Cell{}
	mult := int(math.Sqrt(float64(b.Size)))
	for y := ((c.Y / mult) * b.Size); y < mult; y++ {
		for x := ((c.X / mult) * b.Size); x < mult; x++ {
			result = append(result, b.Cells[b.getIndex(x, y)])
		}
	}
	return result
}
