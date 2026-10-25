// Lời giải: Images
package main

import (
	"image"
	"image/color"

	"tour/internal/pic"
)

type Image struct {
	w, h int
}

// ColorModel cho fmt/image biết cách diễn giải màu do At trả về.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds trả về vùng hợp lệ của ảnh.
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

// At trả về màu tại tọa độ (x, y) — chính là hàm Pic của bài trước,
// chỉ khác là trả về color.Color thay vì uint8.
func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{R: v, G: v, B: 255, A: 255}
}

func main() {
	m := Image{w: 256, h: 256}
	pic.ShowImage(m)
}
