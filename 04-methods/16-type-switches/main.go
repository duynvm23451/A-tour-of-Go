// Type switches / switch theo kiểu
package main

import "fmt"

// Type switch là chuỗi type assertion viết dưới dạng switch.
// Cú pháp giống switch thường, nhưng các case là các KIỂU,
// và biểu thức có dạng  v := i.(type)
func do(i any) {
	switch v := i.(type) {
	case int:
		// Trong nhánh này, v có kiểu int.
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		// Trong nhánh này, v có kiểu string.
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		// Ở default, v giữ nguyên kiểu như i.
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
