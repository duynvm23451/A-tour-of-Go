// Buffered channels / Channel có bộ đệm
package main

import "fmt"

// Channel có thể có bộ đệm: truyền sức chứa làm đối số thứ hai của make.
//
//	ch := make(chan int, 100)
//
// Gửi vào channel có đệm chỉ bị chặn khi bộ đệm ĐẦY.
// Nhận chỉ bị chặn khi bộ đệm RỖNG.
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Bỏ comment dòng dưới sẽ gây deadlock: bộ đệm đầy mà không ai nhận.
	// ch <- 1
	// ch <- 2
	// ch <- 3  // fatal error: all goroutines are asleep - deadlock!
}
