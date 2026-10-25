// Exercise: Readers / Bài tập: io.Reader
//
// ĐỀ BÀI
// Cài đặt một kiểu Reader phát ra dòng ký tự 'A' VÔ HẠN.
//
// Chạy:  go run ./04-methods/22-exercise-reader
// In ra "OK" là đạt.
package main

import "tour/internal/reader"

type MyReader struct{}

// TODO: cài đặt Read sao cho nó điền toàn ký tự 'A' vào b
// và trả về số byte đã điền, err = nil.
func (r MyReader) Read(b []byte) (int, error) {
	return 0, nil
}

func main() {
	reader.Validate(MyReader{})
}
