// If and else / if và else
package main

import (
	"fmt"
	"math"
)

// Biến khai báo trong câu lệnh ngắn của if cũng dùng được trong khối else.
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}

// Ghi chú phong cách: trong code Go thật, người ta thường tránh else sau return.
// Cách viết được ưa chuộng là "early return":
//
//	v := math.Pow(x, n)
//	if v < lim {
//		return v
//	}
//	fmt.Printf("%g >= %g\n", v, lim)
//	return lim
