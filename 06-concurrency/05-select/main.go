// Select
package main

import "fmt"

// The select statement lets a goroutine wait on multiple communication operations.
// select cho phép một goroutine chờ NHIỀU thao tác channel cùng lúc.
//
// select chặn cho tới khi một trong các case sẵn sàng chạy.
// Nếu nhiều case cùng sẵn sàng, nó chọn NGẪU NHIÊN một case.
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
