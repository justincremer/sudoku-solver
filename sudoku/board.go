package sudoku

import (
	"fmt"
	"log"
)

// A board object describing props and containing a list of mutable cells
type Board struct {
	Size       int
	Difficulty int
	Cells      []Cell `json:"squares"`
}

// A single cell in a sudoku board
type Cell struct {
	X   int `json:"x"`
	Y   int `json:"y"`
	Val int `json:"value"`
}

// Initializes a random board with of size 4 or 9 and of difficulty 1, 2, or 3, returning a pointer to said board
func Create(size, difficulty int) *Board {
	board := Board{Size: size, Difficulty: difficulty, Cells: []Cell{}}
	board.zero()
	dto := mapData(&board)
	board.fill(dto)

	return &board
}

// Display's the board in standard form
func (b *Board) Display() {
	fmt.Printf("Size: %v Difficulty: %v\n", b.Size, b.Difficulty)
	s := int(b.Size)
	for i := 0; i < s; i++ {
		fmt.Printf("[%v]", i)
		for j := 0; j < s; j++ {
			c := b.Cells[i*s+j].Val
			fmt.Printf(" %v", c)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// Display's the board as a list of cells
func (b *Board) DisplayChart() {
	fmt.Printf("Size: %v Difficulty: %v\n\n", b.Size, b.Difficulty)
	for i := range b.Cells {
		c := b.Cells[i]
		fmt.Printf("{ X: %v, Y: %v, Value: %v }\n", c.X, c.Y, c.Val)
	}
}

// Zeroes out a board
func (b *Board) zero() {
	cells := []Cell{}
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			cells = append(cells, Cell{X: i, Y: j, Val: 0})
		}
	}

	b.Cells = cells
}

// Partially fills a board with cells in a dataReponse
func (b *Board) fill(dto *Response) {
	for i := range dto.Squares {
		curr := dto.Squares[i]
		index := b.getIndex(curr.X, curr.Y)
		b.Cells[index].Val = curr.Val
	}
}

// Returns a one dimensional index from two dimensional coordinates
func (b *Board) getIndex(x, y int) int {
	return (x * b.Size) + y
}

// Returns two dimensional coordinates from a one dimentional array index
func (b *Board) getCoords(i int) (x, y int) {
	x = i / b.Size
	y = i % b.Size
	return x, y
}

// Handles errors
func must(err error) {
	if err != nil {
		log.Fatalf("Error, %v\n", err)
	}
}
