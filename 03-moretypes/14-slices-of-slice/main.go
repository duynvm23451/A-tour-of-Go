// Slices of slices / Slice của slice
package main

import (
	"fmt"
	"strings"
)

// Slice có thể chứa bất kỳ kiểu nào, kể cả slice khác.
func main() {
	// Tạo bàn cờ tic-tac-toe.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
