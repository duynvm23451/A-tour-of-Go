// Exercise: Errors / Bài tập: Lỗi
//
// ĐỀ BÀI
// Sửa hàm Sqrt (từ bài tập chương trước) để nó trả về giá trị error.
//
// Sqrt phải trả về lỗi khác nil khi nhận số âm, vì nó không hỗ trợ số phức.
// Tạo kiểu mới:
//
//	type ErrNegativeSqrt float64
//
// và cho nó thành error bằng cách cài method:
//
//	func (e ErrNegativeSqrt) Error() string
//
// sao cho ErrNegativeSqrt(-2).Error() trả về
// "cannot Sqrt negative number: -2".
//
// LƯU Ý: gọi fmt.Sprint(e) BÊN TRONG method Error sẽ khiến chương trình
// lặp vô hạn (Sprint gọi lại Error). Hãy ép kiểu trước: fmt.Sprint(float64(e)).
//
// Chạy:  go run ./04-methods/20-exercise-errors
// Test:  go test ./04-methods/20-exercise-errors
package main

import "fmt"

type ErrNegativeSqrt float64

// TODO: cài method Error() string cho ErrNegativeSqrt.

func Sqrt(x float64) (float64, error) {
	// TODO: trả về ErrNegativeSqrt nếu x < 0, ngược lại tính căn bậc hai.
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
