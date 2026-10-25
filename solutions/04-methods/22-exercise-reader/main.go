// Lời giải: Readers (MyReader)
package main

import "tour/internal/reader"

type MyReader struct{}

// Read điền toàn bộ buffer bằng ký tự 'A' rồi báo đã ghi len(b) byte.
//
// Vì dòng dữ liệu là vô hạn nên không bao giờ trả về io.EOF.
// Dùng value receiver để MyReader{} (không phải &MyReader{}) là io.Reader.
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
