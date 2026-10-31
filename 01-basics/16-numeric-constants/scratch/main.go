package main

import (
	"fmt"
)

func main() {
	const Huge = 1 << 70
	fmt.Println(Huge >> 60)

	// var a float64 = Huge ok
	// var b int64 = Huge lỗi
	// var c int64 = Huge >> 10 ok
	const (
		KB = 1 << 10 // 1024
		MB = 1 << 20
		GB = 1 << 30
	)

	fmt.Println("file 3 GB = ", 3*GB/MB, "MB")
}
