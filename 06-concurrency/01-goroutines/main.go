// Goroutines
package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread managed by the Go runtime.
// Goroutine là luồng nhẹ do runtime của Go quản lý.
//
//	go f(x, y, z)
//
// khởi chạy một goroutine mới chạy f(x, y, z).
// Việc TÍNH các đối số x, y, z diễn ra ở goroutine HIỆN TẠI,
// còn f thì chạy ở goroutine MỚI.
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")

	// LƯU Ý: khi hàm main kết thúc, chương trình thoát NGAY,
	// không chờ goroutine nào cả. Ở đây say("hello") chạy đủ lâu
	// nên ta kịp thấy output của goroutine kia.
	// Trong code thật, hãy đồng bộ bằng channel hoặc sync.WaitGroup,
	// KHÔNG dùng time.Sleep.
}
