package main

import (
	"fmt"
	"sudoku-solver/sudoku"
	"time"
)

const (
	size       int = 9
	difficulty int = 2
	tests      int = 100
)

func main() {
	writer := make(chan string)
	defer close(writer)

	singleTest(size, difficulty)
	fmt.Printf("\n")

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
	func() {
		for i := 0; i < tests; i++ {
			go func(i int) {
				board := sudoku.Create(size, difficulty)
				start := time.Now().Nanosecond()
				board.Solve()
				elapsed := ((time.Now().Nanosecond() - start) / 1000000)
				writer <- fmt.Sprintf("Board %d: %d steps, %dms", i+1, board.Steps, elapsed)
			}(i)
		}
	}()
}
