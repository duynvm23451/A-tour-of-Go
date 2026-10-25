// Interface values with nil underlying values / Interface chứa giá trị nil
package main

import "fmt"

// Nếu giá trị cụ thể bên trong interface là nil, method vẫn được GỌI BÌNH THƯỜNG,
// với receiver là nil. Nó KHÔNG gây panic.
//
// Trong Go, viết method chịu được receiver nil là chuyện hoàn toàn hợp lệ
// (xem method String của *tree.Tree trong repo này).
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i) // (<nil>, *main.T) — interface KHÔNG nil, kiểu cụ thể là *T
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
