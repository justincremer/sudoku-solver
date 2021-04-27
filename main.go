package main

import (
	"fmt"
	"sudoku-solver/sudoku"
	"time"
)

const (
	size       int = 9
	difficulty int = 1
	tests      int = 10
)

func main() {
	writer := make(chan string)
	defer close(writer)

	singleTest(size, difficulty)
	fmt.Printf("\n")

	go multiTest(writer)
	for i := range writer {
		fmt.Printf("%v\n", i)
	}
}

func singleTest(size, difficulty int) {
	board := sudoku.Create(9, 3)
	board.Display()
	board.Solve()
	board.Display()
}

func multiTest(writer chan<- string) {
	for i := 1; i <= tests; i++ {
		go func(i int) {
			board := sudoku.Create(size, difficulty)
			start := time.Now().Nanosecond()
			board.Solve()
			elapsed := ((time.Now().Nanosecond() - start) / 1000000)
			writer <- fmt.Sprintf("Board %d: %d steps, %dms", i, board.Steps, elapsed)
		}(i)
	}
}
