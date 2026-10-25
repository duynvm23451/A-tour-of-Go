// Pointers to structs / Con trỏ tới struct
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v

	// Đúng ra phải viết (*p).X, nhưng cách đó rườm rà nên Go cho phép
	// viết tắt là p.X — trình biên dịch tự dereference giùm.
	p.X = 1e9
	fmt.Println(v)
}
