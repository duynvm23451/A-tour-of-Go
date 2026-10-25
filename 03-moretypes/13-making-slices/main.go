// Creating a slice with make / Tạo slice bằng make
package main

import "fmt"

// make cấp phát một mảng đã được zero rồi trả về slice trỏ vào mảng đó.
//
//	make([]T, len)
//	make([]T, len, cap)
func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
