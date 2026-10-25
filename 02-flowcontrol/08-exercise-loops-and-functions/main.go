// Exercise: Loops and Functions / Bài tập: Vòng lặp và Hàm
//
// ĐỀ BÀI
// Cài đặt hàm căn bậc hai bằng phương pháp Newton.
// Bắt đầu với một phỏng đoán z bất kỳ (ví dụ z = 1), rồi lặp lại:
//
//	z -= (z*z - x) / (2*z)
//
// Mỗi vòng lặp z sẽ tiến gần hơn tới căn bậc hai của x.
//
// Yêu cầu nâng dần:
//  1. Lặp 10 lần rồi in kết quả, so sánh với math.Sqrt.
//  2. Đổi điều kiện dừng: lặp cho tới khi z gần như không đổi nữa
//     (thay đổi nhỏ hơn một epsilon rất nhỏ). Nhiều hay ít hơn 10 vòng?
//  3. Thử các giá trị khởi tạo z khác nhau xem hội tụ nhanh cỡ nào.
//
// Chạy thử:  go run ./02-flowcontrol/08-exercise-loops-and-functions
// Chấm bài:  go test ./02-flowcontrol/08-exercise-loops-and-functions
package main

import (
	"fmt"
)

// Sqrt trả về căn bậc hai của x.
func Sqrt(x float64) float64 {
	// TODO: cài đặt bằng phương pháp Newton.
	return 0
}

func main() {
	for _, x := range []float64{2, 4, 9, 16, 100} {
		fmt.Printf("Sqrt(%v) = %v\n", x, Sqrt(x))
	}
}
