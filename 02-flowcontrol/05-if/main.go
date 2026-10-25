// If / Câu lệnh if
package main

import (
	"fmt"
	"math"
)

// Giống for: KHÔNG cần ngoặc tròn quanh điều kiện, nhưng { } thì bắt buộc.
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
