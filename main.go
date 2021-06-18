package main

import (
	"fmt"
	"sudoku-solver/sudoku"
	"time"
)

var (
	size       int
	difficulty int
	tests      int
)

func init() {
	size = 9
	difficulty = 2
	tests = 50
}

func main() {
	writer := make(chan string)
	defer close(writer)

	// singleTest(size, difficulty)
	// fmt.Printf("\n")

	go multiTest(writer)
	for i := 0; i < tests; i++ {
		fmt.Println(<-writer)
	}
}

func singleTest(size, difficulty int) {
	board := sudoku.Create(9, 2)
	board.Display()
	board.Solve()
	board.Display()
}

func multiTest(writer chan<- string) {
	for i := 0; i < tests; i++ {
		go func(i int) {
			board := sudoku.Create(size, difficulty)
			start := time.Now()
			board.Solve()
			elapsed := time.Since(start) / 1000000
			writer <- fmt.Sprintf("Board %d: %d steps, %dms", i+1, board.Steps, elapsed)
		}(i)
	}
}
