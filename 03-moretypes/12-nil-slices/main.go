// Nil slices / Slice nil
package main

import "fmt"

// The zero value of a slice is nil: length 0, capacity 0, no underlying array.
// Zero value của slice là nil: len 0, cap 0, không có mảng nền.
//
// Điểm hay: nil slice DÙNG ĐƯỢC NGAY. append vào nil slice hoàn toàn hợp lệ,
// range trên nil slice cũng không panic. Không cần khởi tạo trước.
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	s = append(s, 1) // hợp lệ
	fmt.Println(s)
}
