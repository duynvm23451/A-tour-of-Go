// Pointers and functions / Con trỏ với hàm thường
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Viết lại Scale và Abs thành HÀM THƯỜNG để thấy sự khác biệt.
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	// Với HÀM, bạn buộc phải truyền con trỏ tường minh: &v.
	// Bỏ dấu & đi là lỗi biên dịch.
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
