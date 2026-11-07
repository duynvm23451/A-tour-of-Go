package main

import (
	"fmt"
	"strings"
)

func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
	}
}

func makeGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)
	for i := 0; i< rows; i++ {
		grid[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			grid[i][j] = i*j
		}
	}
	return grid
}

func main() {
	board := [][]string{
		{"X", "_", "_"},
		{"_", "X", "_"},
		{"_", "_", "X"},
	}

	printBoard(board)

	grid := makeGrid(3,4)
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}

	// g := make([][]int, 2)
	// g[0][0] = 1 // error vì mới chỉ khai báo slice ngoài các slice bên trong đang là nil không thể gán phần tử
}