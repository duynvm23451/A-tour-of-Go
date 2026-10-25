// Forever / Lặp vô hạn
package main

import "fmt"

func main() {
	// Bỏ luôn điều kiện thì vòng lặp chạy mãi mãi.
	// Đây là mẫu rất phổ biến trong Go (vòng lặp sự kiện, server loop...).
	count := 0
	for {
		count++
		if count == 3 {
			// break thoát khỏi vòng lặp; continue nhảy sang vòng kế tiếp.
			break
		}
	}
	fmt.Println("thoát sau", count, "vòng lặp")

	// Bỏ comment dòng dưới để thấy vòng lặp thật sự vô hạn (Ctrl+C để dừng):
	// for {
	// }
}
