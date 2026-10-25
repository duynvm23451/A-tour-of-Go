// Functions / Hàm
package main

import "fmt"

// Notice the type comes AFTER the variable name (unlike C/Java).
// Chú ý: kiểu đứng SAU tên biến (ngược với C/Java).
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
