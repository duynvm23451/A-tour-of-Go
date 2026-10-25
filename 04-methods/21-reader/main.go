// Readers / io.Reader
package main

import (
	"fmt"
	"io"
	"strings"
)

// Package io định nghĩa interface io.Reader — trừu tượng cho "đầu đọc dữ liệu":
//
//	func (T) Read(b []byte) (n int, err error)
//
// Read điền dữ liệu vào slice b, trả về số byte đã điền và lỗi.
// Trả về io.EOF khi hết dữ liệu.
//
// io.Reader được cài đặt ở khắp stdlib: file, kết nối mạng, bộ nén, mật mã...
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
