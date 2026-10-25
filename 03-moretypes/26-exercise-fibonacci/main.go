// Exercise: Fibonacci closure / Bài tập: Closure Fibonacci
//
// ĐỀ BÀI
// Cài đặt hàm fibonacci: nó trả về MỘT HÀM (closure) mà mỗi lần gọi sẽ
// trả về số Fibonacci kế tiếp: 0, 1, 1, 2, 3, 5, 8, 13, ...
//
// Gợi ý: closure cần nhớ hai số gần nhất giữa các lần gọi.
//
// Chạy:  go run ./03-moretypes/26-exercise-fibonacci
// Test:  go test ./03-moretypes/26-exercise-fibonacci
package main

import "fmt"

// fibonacci trả về một hàm sinh lần lượt các số Fibonacci.
func fibonacci() func() int {
	// TODO: khai báo biến trạng thái ở đây, rồi trả về closure dùng chúng.
	return func() int {
		return 0
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
