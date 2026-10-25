// Choosing a value or pointer receiver / Chọn value hay pointer receiver
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// HAI LÝ DO để dùng pointer receiver:
//  1. Method cần SỬA giá trị mà receiver trỏ tới.
//  2. Tránh COPY toàn bộ giá trị ở mỗi lần gọi (hiệu quả hơn với struct lớn).
//
// QUY TẮC THỰC TẾ: một kiểu nên NHẤT QUÁN — hoặc tất cả method dùng
// pointer receiver, hoặc tất cả dùng value receiver. Trộn lẫn dễ gây nhầm.
// Ở đây cả Scale lẫn Abs đều dùng *Vertex dù Abs không cần sửa gì.
func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Trước scale: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("Sau scale: %+v, Abs: %v\n", v, v.Abs())
}
