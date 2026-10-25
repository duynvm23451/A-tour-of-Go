// Channels / Kênh
package main

import "fmt"

// Channels are a typed conduit through which you can send and receive values
// with the channel operator <-.
//
// Channel là ống dẫn có kiểu, dùng toán tử <- để gửi và nhận:
//
//	ch <- v    // gửi v vào channel ch
//	v := <-ch  // nhận từ ch và gán cho v
//
// Mũi tên luôn chỉ hướng dòng dữ liệu chảy.
//
// Channel phải được tạo trước khi dùng: ch := make(chan int)
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // gửi tổng vào c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // nhận hai kết quả

	fmt.Println(x, y, x+y)

	// Mặc định việc gửi và nhận là BLOCKING cho tới khi phía kia sẵn sàng.
	// Nhờ đó goroutine tự đồng bộ với nhau mà không cần khóa tường minh.
}
