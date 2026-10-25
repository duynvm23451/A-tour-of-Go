// Struct fields / Truy cập trường của struct
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4 // truy cập bằng dấu chấm
	fmt.Println(v.X)
}
