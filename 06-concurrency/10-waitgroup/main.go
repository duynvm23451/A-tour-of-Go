// Bonus: sync.WaitGroup
//
// Không có trong Tour of Go gốc nhưng là cách chuẩn để CHỜ goroutine xong.
// Trong code thật, đây là thứ bạn dùng thay cho time.Sleep.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	results := make([]int, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1) // đăng ký thêm một việc cần chờ
		go func() {
			defer wg.Done() // báo xong, luôn đặt trong defer
			results[i] = i * i
		}()
	}

	wg.Wait() // chặn cho tới khi mọi Done đã được gọi
	fmt.Println(results)

	// Ghi chú về biến vòng lặp: từ Go 1.22, mỗi vòng lặp có biến i RIÊNG,
	// nên closure ở trên bắt đúng giá trị. Trước 1.22 phải viết
	// go func(i int) { ... }(i) nếu không cả 5 goroutine sẽ thấy cùng một i.
	//
	// Từ Go 1.25 còn có wg.Go(func(){ ... }) gộp Add/Done vào một lời gọi.
}
