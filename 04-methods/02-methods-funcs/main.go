// Methods are functions / Method chỉ là hàm
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Cùng một logic, viết dạng hàm thường: không có gì khác biệt về mặt bản chất.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
