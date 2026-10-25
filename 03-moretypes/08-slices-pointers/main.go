// Slices are like references to arrays / Slice giống tham chiếu tới mảng
package main

import "fmt"

// A slice does not store any data, it just describes a section of an underlying array.
// Slice KHÔNG lưu dữ liệu, nó chỉ mô tả một đoạn của mảng nền (backing array).
// Sửa phần tử của slice sẽ sửa luôn mảng nền, và mọi slice khác cùng chia sẻ
// mảng nền đó cũng thấy thay đổi.
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b) // a cũng thấy XXX
	fmt.Println(names)
}
