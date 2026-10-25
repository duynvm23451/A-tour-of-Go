// Lời giải: Errors (ErrNegativeSqrt)
package main

import "fmt"

type ErrNegativeSqrt float64

// Error cài đặt interface error.
//
// BẪY: nếu viết fmt.Sprint(e) ở đây, fmt sẽ thấy e có method Error
// và gọi lại chính method này -> đệ quy vô hạn, tràn stack.
// Phải ép về float64 trước để cắt đứt vòng lặp đó.
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		// Quy ước Go: khi có lỗi, giá trị còn lại trả về zero value.
		return 0, ErrNegativeSqrt(x)
	}

	if x == 0 {
		return 0, nil
	}

	z := x
	for i := 0; i < 100; i++ {
		next := z - (z*z-x)/(2*z)
		if next == z {
			break
		}
		z = next
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
