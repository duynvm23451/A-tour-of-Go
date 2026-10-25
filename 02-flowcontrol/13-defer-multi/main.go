// Stacking defers / Xếp chồng defer
package main

import "fmt"

// Deferred calls are pushed onto a stack: they run in LAST-IN-FIRST-OUT order.
// Các lời gọi defer được đẩy vào một ngăn xếp, chạy theo thứ tự VÀO SAU RA TRƯỚC.
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	// Output sẽ đếm ngược 9, 8, 7, ..., 0
}
