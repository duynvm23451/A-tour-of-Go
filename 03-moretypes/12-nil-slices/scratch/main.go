package main

import "fmt"

func isEmpty(s []string) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func main() {
	var a []string
	b := []string{}
	fmt.Println(len(a))   // 0
	fmt.Println(len(b))   // 0
	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false
	fmt.Printf("%#v", a)  // []string(nil)
	fmt.Printf("%#v", b)  // []string{}

	fmt.Println(isEmpty(a))
	fmt.Println(isEmpty(b))

	var s []int
	s = s[:0]
	// vẫn là nil vì không có array nền
	fmt.Println(s == nil)
}
