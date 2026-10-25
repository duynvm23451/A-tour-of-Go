// Methods / Phương thức
package main

import (
	"fmt"
	"math"
)

// Go does NOT have classes. However, you can define methods on types.
// Go KHÔNG có class. Nhưng bạn có thể định nghĩa method trên một kiểu.
//
// Method chỉ là hàm có thêm một tham số đặc biệt: RECEIVER,
// nằm trong cặp ngoặc riêng ngay trước tên hàm.
type Vertex struct {
	X, Y float64
}

// v là receiver.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
