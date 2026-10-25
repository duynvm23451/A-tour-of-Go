// Defer / Hoãn thi hành
package main

import "fmt"

// A defer statement defers the execution of a function until the surrounding
// function returns. The deferred call's ARGUMENTS are evaluated immediately.
//
// defer hoãn việc gọi hàm cho tới khi hàm bao quanh nó return.
// Nhưng THAM SỐ của lời gọi được tính NGAY tại thời điểm gặp defer.
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// defer dùng để dọn dẹp tài nguyên, ví dụ:
//
//	f, err := os.Open(name)
//	if err != nil {
//		return err
//	}
//	defer f.Close()
