package main

import "fmt"

type Cell struct {
	adjacentBombCount int
	bomb bool
	covered bool
	flagged bool
}

func InitializeBoard(rows int, cols int) [][]Cell {
	board := make([][]Cell, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]Cell, cols)
		for j := 0; j < cols; j++ {
			board[i][j] = *new(Cell)
		}
	}
	return board
}

func PrintHorizontalLine(size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
}

func PrintRow(rowNum int, row []Cell) {
	size := len(row)
	lineSize := len(row) * 2 + len(row) + 1

	if rowNum == 0 {
		PrintHorizontalLine(lineSize)
	}

	for i := 0; i < size; i++ {
		if i == 0 {
			fmt.Printf("|")
		}
		fmt.Printf("%x%x|", rowNum, i)
	}
	fmt.Printf("\n")

	for i := 0; i < size; i++ {
		if i == 0 {
			fmt.Printf("|")
		}
		fmt.Printf("  |")
	}
	fmt.Printf("\n")

	PrintHorizontalLine(lineSize)
}

func PrintBoard(board [][]Cell) {
	for i := 0; i < len(board); i++ {
		PrintRow(i, board[i])
	}
}


func main() {
	board := InitializeBoard(16, 16)
	PrintBoard(board)
}
