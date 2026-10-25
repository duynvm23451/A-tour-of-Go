// Bonus: context.Context
//
// Không có trong Tour of Go gốc, nhưng bắt buộc phải biết nếu viết server.
// context truyền tín hiệu HỦY và HẠN CHÓT xuyên qua các lời gọi hàm.
package main

import (
	"context"
	"fmt"
	"time"
)

// Quy ước: context là THAM SỐ ĐẦU TIÊN và luôn đặt tên là ctx.
func work(ctx context.Context, name string, d time.Duration) error {
	select {
	case <-time.After(d):
		fmt.Println(name, "hoàn thành")
		return nil
	case <-ctx.Done():
		// ctx.Done() đóng lại khi bị hủy hoặc quá hạn.
		fmt.Println(name, "bị hủy:", ctx.Err())
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	// defer cancel() luôn đi kèm WithTimeout/WithCancel để giải phóng tài nguyên.
	defer cancel()

	work(ctx, "việc nhanh", 50*time.Millisecond)
	work(ctx, "việc chậm", 500*time.Millisecond)
}
