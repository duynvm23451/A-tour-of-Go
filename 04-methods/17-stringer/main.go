// Stringers / Interface Stringer
package main

import "fmt"

// Một trong những interface phổ biến nhất, nằm ở package fmt:
//
//	type Stringer interface {
//		String() string
//	}
//
// Các hàm của fmt sẽ tự gọi String() để in giá trị.
// Cài Stringer cho kiểu của bạn là cách chuẩn để kiểm soát cách nó hiển thị.
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
