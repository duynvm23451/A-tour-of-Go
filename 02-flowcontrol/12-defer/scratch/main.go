package main

import (
	"fmt"
	"time"
)

func process(n int) string {
	defer fmt.Println("dọn dẹp xong")
	if n < 0 {
		return "lỗi" // return sớm, defer vẫn chạy
	}
	return "ok" // return bình thường, defer cũng chạy
}

func slow() {
	start := time.Now()
	defer fmt.Println("mất", time.Since(start))
	time.Sleep(500 * time.Millisecond)
}

func main() {
	fmt.Println(process(-1))
	fmt.Println(process(5))
	// dọn dẹp xong
	// lỗI
	// dọn dẹp xong
	// ok

	for i := 0; i < 3; i++ {
		defer fmt.Println("defer", i)
	}
	fmt.Println("hết vòng lặp")

	// hết vòng lặp
	// defer 2
	// defer 1
	// defer 0
}
