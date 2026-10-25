// sync.Mutex
package main

import (
	"fmt"
	"sync"
	"time"
)

// Channel dùng để GIAO TIẾP giữa các goroutine.
// Nhưng khi chỉ cần đảm bảo mỗi lúc chỉ một goroutine chạm vào một biến
// (mutual exclusion / loại trừ lẫn nhau), hãy dùng sync.Mutex.
//
// SafeCounter là map an toàn khi dùng đồng thời.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc tăng bộ đếm của key đã cho.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Khóa để mỗi lúc chỉ một goroutine ghi vào map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value trả về giá trị hiện tại của key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// defer đảm bảo mở khóa kể cả khi hàm return sớm hoặc panic.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))

	// Thử chạy với cờ phát hiện tranh chấp dữ liệu:
	//	go run -race ./06-concurrency/08-mutex-counter
	// Bỏ Lock/Unlock đi rồi chạy lại để thấy -race báo lỗi.
}
