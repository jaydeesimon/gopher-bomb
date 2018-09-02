package main

import "fmt"
import "math/rand"
import "time"

const Rows = 16
const Cols = 16
const BombPercentage = 10

type Cell struct {
	adjacentBombCount int
	bomb              bool
	covered           bool
	flagged           bool
}

func TranslateToXY(n int) (int, int) {
	return int(n / Rows), n % Rows
}

func InitializeBombs(board [][]Cell) [][]Cell {
	boardSize := float64(Rows * Cols)
	bombCount := int(float64(boardSize * (BombPercentage / 100.0)))

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i, v := range r.Perm(Rows * Cols) {
		if i >= bombCount {
			break
		}

		x, y := TranslateToXY(v)
		board[x][y].bomb = true
	}

	return board
}

func InitializeBoard(rows int, cols int) [][]Cell {
	board := make([][]Cell, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]Cell, cols)
		for j := 0; j < cols; j++ {
			cell := *new(Cell)
			cell.covered = true
			board[i][j] = cell
		}
	}

	board = InitializeBombs(board)

	return board
}

func PrintHorizontalLine(size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
}

func CellFormat(cell Cell) string {
	if cell.bomb {
		return "☣️"
	}

	return "#️⃣"
}

func PrintRow(rowNum int, row []Cell) {
	size := len(row)
	lineSize := len(row)*2 + len(row) + 1

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
		fmt.Printf("%s |", CellFormat(row[i]))
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
	board := InitializeBoard(Rows, Cols)
	PrintBoard(board)
}
