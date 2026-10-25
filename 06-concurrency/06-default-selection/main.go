// Default Selection / Nhánh default của select
package main

import (
	"fmt"
	"time"
)

// Nhánh default chạy khi KHÔNG case nào khác sẵn sàng.
// Dùng nó để gửi/nhận mà không bị chặn (non-blocking):
//
//	select {
//	case i := <-c:
//		// dùng i
//	default:
//		// nhận từ c sẽ bị chặn, nên làm việc khác
//	}
func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
