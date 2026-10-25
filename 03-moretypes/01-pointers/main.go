// Pointers / Con trỏ
package main

import "fmt"

// Go có con trỏ. Con trỏ giữ ĐỊA CHỈ BỘ NHỚ của một giá trị.
// Kiểu *T là con trỏ tới giá trị kiểu T. Zero value của nó là nil.
//
// KHÁC C: Go KHÔNG có phép toán số học trên con trỏ (pointer arithmetic).
func main() {
	i, j := 42, 2701

	p := &i         // & lấy địa chỉ của i -> p có kiểu *int
	fmt.Println(*p) // * đọc giá trị qua con trỏ ("dereference")
	*p = 21         // ghi giá trị qua con trỏ
	fmt.Println(i)  // i đã đổi thành 21

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
