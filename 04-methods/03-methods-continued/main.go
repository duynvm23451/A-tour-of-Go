// Methods continued / Method (tiếp)
package main

import (
	"fmt"
	"math"
)

// Method khai báo được trên MỌI kiểu có tên (named type), không riêng struct.
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// GIỚI HẠN: chỉ khai báo được method trên kiểu định nghĩa TRONG CÙNG PACKAGE.
	// Không thể thêm method cho int, hay cho kiểu của package khác.
	// Muốn vậy phải tạo kiểu mới bọc lại, như MyFloat ở trên.
}
