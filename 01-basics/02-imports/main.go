// Imports / Nhập gói
package main

// This is a "factored" import statement — the idiomatic style.
// Đây là cách import gộp khối — kiểu viết idiomatic của Go.
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
