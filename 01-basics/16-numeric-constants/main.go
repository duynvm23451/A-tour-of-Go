// Numeric constants / Hằng số dạng số
package main

import "fmt"

// Numeric constants are high-precision values with no fixed type.
// Hằng số dạng số là GIÁ TRỊ có độ chính xác cao và KHÔNG mang kiểu cố định.
// Chúng chỉ nhận kiểu khi được dùng trong một ngữ cảnh cụ thể.
const (
	// Dịch trái 1 đi 100 bit: lớn hơn mọi kiểu int của Go.
	Big = 1 << 100
	// Dịch phải 99 bit -> còn 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// needInt(Big) sẽ KHÔNG biên dịch được: tràn kiểu int.
	// fmt.Println(needInt(Big))
}
