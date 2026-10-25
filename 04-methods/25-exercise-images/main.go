// Exercise: Images / Bài tập: Ảnh
//
// ĐỀ BÀI
// Nhớ lại bài tập vẽ ảnh trước đây. Lần này hãy trả về một cài đặt
// của interface image.Image thay vì slice.
//
// Định nghĩa kiểu Image của riêng bạn, cài đủ các method cần thiết,
// rồi gọi pic.ShowImage.
//
//   - Bounds phải trả về image.Rectangle, ví dụ image.Rect(0, 0, w, h).
//   - ColorModel phải trả về color.RGBAModel.
//   - At trả về một màu; color.RGBA{v, v, 255, 255} cho tông xanh dương.
//
// Chạy:  go run ./04-methods/25-exercise-images
package main

import (
	"image"
	"image/color"

	"tour/internal/pic"
)

type Image struct {
	// TODO: thêm các trường bạn cần, ví dụ w, h int.
}

func (i Image) ColorModel() color.Model {
	// TODO
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	// TODO: trả về kích thước thật của ảnh.
	return image.Rect(0, 0, 256, 256)
}

func (i Image) At(x, y int) color.Color {
	// TODO: tính màu từ x, y.
	return color.RGBA{0, 0, 0, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
