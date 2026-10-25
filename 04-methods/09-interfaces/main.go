// Interfaces / Giao diện
package main

import (
	"fmt"
	"math"
)

// An interface type is a set of method signatures.
// Kiểu interface là một TẬP HỢP CÁC CHỮ KÝ METHOD.
// Một giá trị interface có thể chứa bất kỳ giá trị nào cài đặt đủ các method đó.
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // MyFloat có method Abs với value receiver
	a = &v // *Vertex có method Abs với pointer receiver

	// a = v  // LỖI BIÊN DỊCH!
	// Vertex (value) KHÔNG cài đặt Abser, vì Abs được khai báo trên *Vertex.
	// Method set của Vertex không chứa các method có pointer receiver.
	// Đây là lỗi mà người mới học Go gặp nhiều nhất.

	fmt.Println(a.Abs())
}
