// Package reader is a local stand-in for golang.org/x/tour/reader.
// Package reader là bản thay thế cục bộ cho golang.org/x/tour/reader,
// dùng để chấm bài tập io.Reader.
package reader

import (
	"fmt"
	"io"
	"os"
)

// Validate checks that r behaves like an infinite stream of 'A' bytes.
// Validate kiểm tra r có hành xử như một dòng byte 'A' vô hạn hay không.
func Validate(r io.Reader) {
	// Thử nhiều kích thước buffer khác nhau: một Reader đúng phải hoạt động
	// với mọi kích thước, kể cả buffer rất nhỏ.
	for _, size := range []int{1, 2, 7, 512, 4096} {
		buf := make([]byte, size)
		total := 0
		for total < 8192 {
			n, err := r.Read(buf)
			if err != nil {
				fail("Read(buf[%d]) returned unexpected error: %v", size, err)
				return
			}
			if n <= 0 {
				fail("Read(buf[%d]) returned n = %d; want n > 0", size, n)
				return
			}
			if n > size {
				fail("Read(buf[%d]) returned n = %d; want n <= %d", size, n, size)
				return
			}
			for i := 0; i < n; i++ {
				if buf[i] != 'A' {
					fail("Read produced byte %q at offset %d; want 'A'", buf[i], total+i)
					return
				}
			}
			total += n
		}
	}
	fmt.Println("OK")
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stdout, "FAIL: "+format+"\n", args...)
}
