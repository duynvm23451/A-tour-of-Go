// Type conversions / Ép kiểu
package main

import (
	"fmt"
	"math"
)

func main() {
	// The expression T(v) converts the value v to the type T.
	// Biểu thức T(v) chuyển giá trị v sang kiểu T.
	//
	// Khác với C, Go KHÔNG tự động ép kiểu giữa các kiểu số:
	// phải viết ép kiểu tường minh, nếu không sẽ lỗi biên dịch.
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// Thử bỏ float64(...) và uint(...) để xem lỗi biên dịch.
}
