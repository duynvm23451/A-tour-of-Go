// Bonus: khi nào tham số của defer được tính?
//
// Đây là chỗ người mới hay nhầm nhất, nên tách hẳn ra một bài.
package main

import "fmt"

func main() {
	i := 0
	defer fmt.Println("defer thấy i =", i) // i được COPY ngay bây giờ -> 0
	i = 99
	fmt.Println("cuối hàm i =", i) // 99

	// Muốn defer nhìn thấy giá trị MỚI NHẤT thì bọc trong closure:
	j := 0
	defer func() {
		fmt.Println("closure thấy j =", j) // đọc j lúc chạy -> 42
	}()
	j = 42
}
