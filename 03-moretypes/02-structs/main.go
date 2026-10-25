// Structs / Kiểu cấu trúc
package main

import "fmt"

// A struct is a collection of fields.
// struct là một tập hợp các trường (field).
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
