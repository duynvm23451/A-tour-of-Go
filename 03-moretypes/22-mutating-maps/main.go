// Mutating maps / Thao tác với map
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42 // thêm hoặc cập nhật
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") // xóa key
	fmt.Println("The value:", m["Answer"])

	// Cú pháp "comma ok": kiểm tra key có tồn tại không.
	// Nếu key không có, v nhận ZERO VALUE và ok = false.
	// Đây là cách duy nhất để phân biệt "không có key" với "có key, giá trị 0".
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// Ghi chú: thứ tự duyệt map là NGẪU NHIÊN theo thiết kế.
	// Muốn output ổn định thì phải lấy key ra rồi sort.
}
