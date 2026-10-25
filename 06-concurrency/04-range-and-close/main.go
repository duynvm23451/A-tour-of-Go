// Range and Close / range và close trên channel
package main

import "fmt"

// Bên GỬI có thể close một channel để báo rằng sẽ không còn giá trị nào nữa.
// Bên NHẬN kiểm tra bằng dạng hai giá trị:
//
//	v, ok := <-ch    // ok == false nghĩa là channel đã đóng và cạn dữ liệu
//
// Vòng lặp  for i := range ch  sẽ nhận liên tục cho tới khi channel đóng.
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	// LƯU Ý:
	//   - Chỉ BÊN GỬI được close, không bao giờ để bên nhận close.
	//     Gửi vào channel đã đóng sẽ panic.
	//   - Không bắt buộc phải close channel. Chỉ close khi bên nhận
	//     cần biết là "hết dữ liệu rồi", ví dụ để thoát vòng range.
}
