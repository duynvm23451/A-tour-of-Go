// Lời giải: Slices (Pic)
package main

import "tour/internal/pic"

func Pic(dx, dy int) [][]uint8 {
	// Bước 1: cấp phát slice NGOÀI, gồm dy dòng.
	img := make([][]uint8, dy)

	for y := range img {
		// Bước 2: cấp phát từng dòng, mỗi dòng dx phần tử.
		row := make([]uint8, dx)
		for x := range row {
			// Công thức tùy chọn. x^y cho hoa văn kiểu fractal rất đẹp.
			// Thử đổi thành (x+y)/2 hoặc x*y để so sánh.
			row[x] = uint8(x ^ y)
		}
		img[y] = row
	}

	return img
}

func main() {
	pic.Show(Pic)
}
