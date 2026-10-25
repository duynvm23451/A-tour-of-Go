// Type assertions / Khẳng định kiểu
package main

import "fmt"

func main() {
	var i any = "hello"

	// Dạng 1: t := i.(T) — nếu i không chứa T thì PANIC.
	s := i.(string)
	fmt.Println(s)

	// Dạng 2: t, ok := i.(T) — dạng "comma ok", KHÔNG panic.
	// Nếu không khớp: t nhận zero value, ok = false.
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// Dòng dưới sẽ panic vì i chứa string chứ không phải float64.
	// f = i.(float64)
	// fmt.Println(f)
}
