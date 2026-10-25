// Methods and pointer indirection (2) / Chiều ngược lại
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Value receiver.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	// Method với value receiver gọi được trên con trỏ:
	// p.Abs() được hiểu là (*p).Abs().
	fmt.Println(p.Abs())

	// Nhưng HÀM nhận Vertex thì phải dereference tường minh.
	fmt.Println(AbsFunc(*p))
}
