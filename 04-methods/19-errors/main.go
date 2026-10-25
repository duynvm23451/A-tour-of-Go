// Errors / Lỗi
package main

import (
	"fmt"
	"time"
)

// Go biểu diễn trạng thái lỗi bằng GIÁ TRỊ kiểu error — một interface có sẵn:
//
//	type error interface {
//		Error() string
//	}
//
// Go KHÔNG có exception. Hàm trả về error, và người gọi kiểm tra nó.
// Quy ước: error == nil nghĩa là thành công.
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// Ghi chú cho code thật:
//   - errors.New("...") và fmt.Errorf("...") để tạo lỗi đơn giản.
//   - fmt.Errorf("đọc config: %w", err) để BỌC lỗi, giữ lại lỗi gốc.
//   - errors.Is(err, ErrNotFound) so sánh xuyên qua các lớp bọc.
//   - errors.As(err, &myErr) lấy ra lỗi thuộc kiểu cụ thể.
