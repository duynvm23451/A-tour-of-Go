// Images / Ảnh
package main

import (
	"fmt"
	"image"
)

// Package image định nghĩa interface Image:
//
//	type Image interface {
//		ColorModel() color.Model
//		Bounds() Rectangle
//		At(x, y int) color.Color
//	}
//
// Lưu ý: Rectangle trả về từ Bounds nằm trong chính package image.
func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
