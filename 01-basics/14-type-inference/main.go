// Type inference / Suy luận kiểu
package main

import "fmt"

func main() {
	// Khi khai báo không ghi kiểu, kiểu được suy ra từ vế phải.
	v := 42 // int

	// Khi vế phải là hằng số không mang kiểu (untyped constant),
	// kiểu suy ra phụ thuộc độ chính xác của hằng số đó:
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}
