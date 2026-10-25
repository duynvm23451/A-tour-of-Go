// Appending to a slice / Thêm phần tử vào slice
package main

import "fmt"

// func append(s []T, vs ...T) []T
//
// Nếu mảng nền còn đủ chỗ (cap), append ghi thẳng vào đó.
// Nếu KHÔNG đủ, Go cấp phát mảng lớn hơn và COPY sang — lúc này slice cũ
// và slice mới không còn chung mảng nền nữa. Đây là nguồn gốc của nhiều bug.
//
// Vì vậy LUÔN gán lại kết quả:  s = append(s, x)
func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	// Thêm nhiều phần tử một lúc.
	s = append(s, 2, 3, 4)
	printSlice(s)

	// Nối hai slice: cần toán tử ... để "trải" slice thành các đối số rời.
	other := []int{5, 6}
	s = append(s, other...)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
