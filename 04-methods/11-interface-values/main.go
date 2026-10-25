// Interface values / Giá trị interface
package main

import (
	"fmt"
	"math"
)

// Có thể hình dung giá trị interface như một cặp (value, type):
//
//	(giá trị cụ thể, kiểu cụ thể)
//
// Gọi method trên interface sẽ chạy method của KIỂU CỤ THỂ bên trong.
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
