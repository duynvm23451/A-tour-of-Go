// Zero values / Giá trị không
package main

import "fmt"

// Variables declared without an explicit initial value are given their zero value.
// Biến khai báo mà không gán giá trị sẽ nhận GIÁ TRỊ KHÔNG (zero value):
//
//	0     cho kiểu số
//	false cho bool
//	""    cho string
//	nil   cho pointer, slice, map, channel, func, interface
//
// Go không có khái niệm "biến chưa khởi tạo" như C — điểm này rất quan trọng.
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
