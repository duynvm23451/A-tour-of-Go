// Slice length and capacity / Độ dài và sức chứa của slice
package main

import "fmt"

// len(s) = số phần tử slice đang chứa.
// cap(s) = số phần tử của mảng nền, tính TỪ phần tử đầu tiên của slice.
//
// Hiểu rõ len vs cap là chìa khóa để không dính bug với append.
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Cắt về rỗng: len = 0 nhưng cap vẫn còn nguyên.
	s = s[:0]
	printSlice(s)

	// Nới lại độ dài — làm được vì cap vẫn đủ.
	s = s[:4]
	printSlice(s)

	// Bỏ hai phần tử đầu: cap giảm theo vì mốc bắt đầu đã dịch.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
