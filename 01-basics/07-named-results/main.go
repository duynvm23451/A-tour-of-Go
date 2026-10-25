// Named return values / Giá trị trả về có tên
package main

import "fmt"

// Return values may be named; they are treated as variables defined at the top
// of the function. A bare "return" is a "naked return".
//
// Giá trị trả về có thể đặt tên; chúng như biến khai báo sẵn ở đầu hàm.
// Lệnh "return" trống sẽ trả về giá trị hiện tại — gọi là "naked return".
//
// Lời khuyên thực tế: chỉ dùng naked return trong hàm NGẮN; hàm dài sẽ khó đọc.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
