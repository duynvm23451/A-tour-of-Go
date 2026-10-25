// Exercise: Maps / Bài tập: Map
//
// ĐỀ BÀI
// Cài đặt WordCount. Hàm này trả về một map chứa số lần xuất hiện của
// mỗi "từ" trong chuỗi s. Hàm wc.Test sẽ chạy thử với vài chuỗi mẫu
// và in PASS/FAIL cho từng trường hợp.
//
// Gợi ý: strings.Fields tách chuỗi thành các từ theo khoảng trắng.
//
// Chạy:  go run ./03-moretypes/23-exercise-maps
// Test:  go test ./03-moretypes/23-exercise-maps
package main

import "tour/internal/wc"

func WordCount(s string) map[string]int {
	// TODO: tạo map, duyệt các từ và đếm.
	return map[string]int{}
}

func main() {
	wc.Test(WordCount)
}
