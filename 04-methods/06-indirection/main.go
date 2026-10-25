// Methods and pointer indirection / Method và việc tự dereference
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// Với METHOD, Go tự chèn &: v.Scale(2) thực chất là (&v).Scale(2).
	v.Scale(2)

	p := &Vertex{4, 3}
	p.Scale(3) // đã là con trỏ, gọi trực tiếp

	// Với HÀM thì KHÔNG có sự tiện lợi đó — phải truyền đúng kiểu.
	ScaleFunc(&v, 10)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
