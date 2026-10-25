// The empty interface / Interface rỗng
package main

import "fmt"

// The interface type that specifies zero methods is the empty interface.
// Interface không có method nào gọi là interface rỗng: interface{}
//
// Nó chứa được GIÁ TRỊ THUỘC BẤT KỲ KIỂU NÀO (mọi kiểu đều có ít nhất 0 method).
// Dùng để xử lý giá trị chưa biết kiểu, ví dụ fmt.Print nhận ...interface{}.
//
// Từ Go 1.18 có alias "any" cho interface{}, và any là cách viết được ưa dùng.
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i any) {
	fmt.Printf("(%v, %T)\n", i, i)
}
