// Exercise: Slices / Bài tập: Slice
//
// ĐỀ BÀI
// Cài đặt hàm Pic. Nó nhận vào dy, dx và trả về một slice chứa dy slice,
// mỗi slice con có dx phần tử kiểu uint8.
//
// Chương trình sẽ đọc mảng đó như các giá trị xám (0 = đen, 255 = trắng)
// và vẽ ra ảnh.
//
// Giá trị mỗi ô do bạn tự chọn. Vài công thức thú vị:
// (x+y)/2, x*y, x^y, x%(y+1), ...
//
// Chạy:  go run ./03-moretypes/18-exercise-slices
// Kết quả in ra terminal dạng ASCII và ghi thêm file out.png.
package main

import "tour/internal/pic"

func Pic(dx, dy int) [][]uint8 {
	// TODO: cấp phát slice ngoài bằng make, rồi cấp phát từng dòng bên trong.
	return nil
}

func main() {
	pic.Show(Pic)
}
