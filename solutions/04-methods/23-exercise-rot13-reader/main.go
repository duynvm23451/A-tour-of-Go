// Lời giải: rot13Reader
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// Read đọc từ Reader bên trong rồi biến đổi TẠI CHỖ đúng n byte đã đọc.
//
// Chú ý chỉ xử lý b[:n]: phần còn lại của buffer là rác từ lần đọc trước.
// Và luôn trả về nguyên vẹn n, err nhận được từ Reader bên trong.
func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func rot13(c byte) byte {
	switch {
	case c >= 'a' && c <= 'z':
		return 'a' + (c-'a'+13)%26
	case c >= 'A' && c <= 'Z':
		return 'A' + (c-'A'+13)%26
	}
	return c // ký tự không phải chữ cái giữ nguyên
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	os.Stdout.WriteString("\n")
}
