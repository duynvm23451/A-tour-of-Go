// Lời giải: Maps (WordCount)
package main

import (
	"strings"

	"tour/internal/wc"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)

	// strings.Fields tách chuỗi theo mọi khoảng trắng và tự bỏ phần rỗng.
	for _, word := range strings.Fields(s) {
		// Không cần kiểm tra key đã tồn tại chưa: đọc key vắng mặt
		// trả về zero value (0), nên counts[word]++ luôn đúng.
		counts[word]++
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
