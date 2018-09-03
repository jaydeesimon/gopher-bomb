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

func SetBombs(board [][]Cell) [][]Cell {
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

func InBounds(x int, y int) bool {
	return x >= 0 && x <= Rows-1 && y >= 0 && y <= Cols-1
}

func CountAdjacentBombs(board [][]Cell, x int, y int) int {
	adjacentBombCount := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			ax := x + i
			ay := y + j

			if InBounds(ax, ay) && board[ax][ay].bomb {
				adjacentBombCount++
			}
		}
	}

	return adjacentBombCount
}

func CountBombs(board [][]Cell) [][]Cell {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j].adjacentBombCount = CountAdjacentBombs(board, i, j)
		}
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

	board = SetBombs(board)
	board = CountBombs(board)

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

	numEmojiMap := map[int]string{
		0: "0️⃣",
		1: "1️⃣",
		2: "2️⃣",
		3: "3️⃣",
		4: "4️⃣",
		5: "5️⃣",
		6: "6️⃣",
		7: "7️⃣",
		8: "8️⃣",
	}

	return numEmojiMap[cell.adjacentBombCount]
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
