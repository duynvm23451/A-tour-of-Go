// Nil interface values / Interface nil thực sự
package main

import "fmt"

type I interface {
	M()
}

func main() {
	// Interface nil thực sự: KHÔNG có giá trị, cũng KHÔNG có kiểu cụ thể.
	// Gọi method trên nó sẽ PANIC ngay lập tức.
	var i I
	describe(i)
	i.M() // panic: runtime error: invalid memory address or nil pointer dereference
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// BẪY KINH ĐIỂN — so sánh bài này với bài 12:
//
//	var t *T = nil
//	var i I = t
//	i == nil  ->  FALSE!
//
// Vì i mang kiểu cụ thể *T, nó không phải interface nil.
// Đây là nguyên nhân hàng đầu của bug "err != nil nhưng err lại là nil".
