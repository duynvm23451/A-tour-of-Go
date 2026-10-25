// Functions continued / Hàm (tiếp)
package main

import "fmt"

// When consecutive parameters share a type, you may omit it from all but the last.
// Khi nhiều tham số liên tiếp cùng kiểu, có thể bỏ kiểu ở tất cả trừ cái cuối:
// (x int, y int) == (x, y int)
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
