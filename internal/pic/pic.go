// Package pic is a local stand-in for golang.org/x/tour/pic.
// Package pic là bản thay thế cục bộ cho golang.org/x/tour/pic,
// nhờ vậy repo này chạy được mà không cần tải module ngoài.
package pic

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
)

const (
	defaultDX = 256
	defaultDY = 256
)

// Show calls f(dx, dy), turns the result into a grayscale image and writes it
// to out.png next to the current working directory.
//
// Show gọi f(dx, dy), đổi kết quả thành ảnh xám rồi ghi ra file out.png.
// Bản gốc trên tour.golang.org hiện ảnh base64 ngay trên trình duyệt;
// ở đây ta ghi ra file + vẽ thêm bản xem nhanh bằng ASCII trong terminal.
func Show(f func(dx, dy int) [][]uint8) {
	ShowImage(gray(f(defaultDX, defaultDY)))
}

// ShowImage writes any image.Image to out.png.
// ShowImage ghi một image.Image bất kỳ ra out.png.
func ShowImage(img image.Image) {
	const name = "out.png"

	f, err := os.Create(name)
	if err != nil {
		fmt.Println("pic: cannot create file:", err)
		return
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		fmt.Println("pic: cannot encode png:", err)
		return
	}

	abs, err := filepath.Abs(name)
	if err != nil {
		abs = name
	}
	fmt.Println("pic: wrote", abs)
	fmt.Println("pic: ASCII preview / bản xem nhanh:")
	printASCII(img)
}

// gray converts the [][]uint8 returned by the exercise into an image.
// gray đổi [][]uint8 mà bài tập trả về thành ảnh xám.
func gray(data [][]uint8) image.Image {
	dy := len(data)
	dx := 0
	if dy > 0 {
		dx = len(data[0])
	}
	if dx == 0 || dy == 0 {
		// Empty result: still return a valid 1x1 image so callers do not panic.
		// Kết quả rỗng: vẫn trả ảnh 1x1 hợp lệ để không panic.
		return image.NewGray(image.Rect(0, 0, 1, 1))
	}

	img := image.NewGray(image.Rect(0, 0, dx, dy))
	for y, row := range data {
		for x, v := range row {
			if x >= dx {
				break
			}
			img.SetGray(x, y, color.Gray{Y: v})
		}
	}
	return img
}

// printASCII downsamples img to a small grid of characters.
// printASCII thu nhỏ ảnh về lưới ký tự để xem ngay trong terminal.
func printASCII(img image.Image) {
	const (
		cols   = 64
		rows   = 32
		ramp   = " .:-=+*#%@"
		levels = len(ramp) - 1
	)

	b := img.Bounds()
	if b.Dx() == 0 || b.Dy() == 0 {
		return
	}

	for r := 0; r < rows; r++ {
		line := make([]byte, cols)
		for c := 0; c < cols; c++ {
			x := b.Min.X + c*b.Dx()/cols
			y := b.Min.Y + r*b.Dy()/rows
			// Luminance 0..65535 -> chỉ số trong ramp.
			lum, _, _, _ := color.Gray16Model.Convert(img.At(x, y)).RGBA()
			line[c] = ramp[int(lum)*levels/0xffff]
		}
		fmt.Println(string(line))
	}
}
