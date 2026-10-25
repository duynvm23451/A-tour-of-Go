// Exercise: rot13Reader / Bài tập: rot13Reader
//
// ĐỀ BÀI
// Một mẫu phổ biến là io.Reader BỌC một io.Reader khác và biến đổi dữ liệu
// trên đường đi (giống gzip.NewReader nhận io.Reader và trả io.Reader).
//
// Hãy cài đặt rot13Reader: nó đọc từ Reader bên trong rồi áp dụng
// mật mã thay thế ROT13 lên mọi ký tự chữ cái.
//
// Kết quả đúng sẽ in ra: "You cracked the code!"
//
// Chạy:  go run ./04-methods/23-exercise-rot13-reader
// Test:  go test ./04-methods/23-exercise-rot13-reader
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// TODO: đọc từ r.r rồi biến đổi ROT13 các byte đọc được trước khi trả về.
func (r *rot13Reader) Read(b []byte) (int, error) {
	return r.r.Read(b)
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	os.Stdout.WriteString("\n")
}
