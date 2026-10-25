// If with a short statement / if kèm câu lệnh khởi tạo
package main

import (
	"fmt"
	"math"
)

// if có thể chạy một câu lệnh ngắn TRƯỚC điều kiện.
// Biến khai báo ở đó chỉ tồn tại trong phạm vi của if — rất idiomatic,
// vì nó giữ biến tạm không rò rỉ ra ngoài.
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// Ở đây không dùng được v: đã ra khỏi scope.
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
