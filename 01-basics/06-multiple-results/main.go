// Multiple results / Trả về nhiều giá trị
package main

import "fmt"

// A function can return any number of results.
// Một hàm có thể trả về bao nhiêu giá trị tùy ý.
// Đây là nền tảng cho cách Go xử lý lỗi: trả về cặp (value, error).
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
