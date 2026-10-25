// Lời giải: Fibonacci closure
package main

import "fmt"

func fibonacci() func() int {
	// Hai biến này sống trong closure: mỗi lần gọi fibonacci()
	// sẽ tạo ra một bộ x, y riêng biệt.
	x, y := 0, 1

	return func() int {
		result := x
		x, y = y, x+y // gán song song, không cần biến tạm
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
