// Pointer receivers / Receiver là con trỏ
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Receiver kiểu *Vertex: method có thể SỬA giá trị mà receiver trỏ tới.
// Đây là lý do chính để dùng pointer receiver.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs()) // 50

	// Nếu Scale dùng value receiver (v Vertex) thì nó chỉ sửa BẢN SAO
	// và v ở đây sẽ không đổi -> kết quả in ra là 5.
}
