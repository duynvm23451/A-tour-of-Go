package main

import "fmt"

func main() {
	// Type inference / Suy luận kiểu
	a := 1e3
	b := 'x' + 1
	c := 7 / 2.0
	var d int64 = 5
	e := d * 2

	fmt.Printf("a is of type %T\n", a) // float64
	fmt.Printf("b is of type %T\n", b) // int32
	fmt.Printf("c is of type %T\n", c) // float64
	fmt.Printf("d is of type %T\n", d) // int64
	fmt.Printf("e is of type %T\n", e) // int64

	total := 0.0
	total = total + 2.5 * 3
	fmt.Println(total)

	x := float64(10)
	y := 10
	fmt.Printf("x is of type %T\n", x) // float64
	fmt.Printf("y is of type %T\n", y) // int
}