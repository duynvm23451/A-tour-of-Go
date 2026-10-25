// Lời giải: Loops and Functions (Newton's method)
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}

	// Phỏng đoán ban đầu: dùng chính x hội tụ nhanh hơn hằng số 1
	// khi x lớn (ví dụ x = 1e6).
	z := x

	// Điều kiện dừng: lặp tới khi z KHÔNG CÒN THAY ĐỔI nữa.
	// Không dùng epsilon cố định như 1e-15 vì với x lớn, sai số làm tròn
	// của float64 vốn đã lớn hơn epsilon đó -> vòng lặp sẽ không bao giờ dừng.
	// Giới hạn 100 vòng để phòng trường hợp giá trị dao động qua lại.
	for i := 0; i < 100; i++ {
		next := z - (z*z-x)/(2*z)
		if next == z {
			break
		}
		z = next
	}
	return z
}

func main() {
	for _, x := range []float64{2, 4, 9, 16, 100} {
		fmt.Printf("Sqrt(%v) = %v   math.Sqrt = %v\n", x, Sqrt(x), math.Sqrt(x))
	}
}
