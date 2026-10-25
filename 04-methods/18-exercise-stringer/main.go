// Exercise: Stringers / Bài tập: Stringer
//
// ĐỀ BÀI
// Cài method String cho kiểu IPAddr để in địa chỉ IP dạng chấm.
// Ví dụ IPAddr{1, 2, 3, 4} phải in ra "1.2.3.4".
//
// Chạy:  go run ./04-methods/18-exercise-stringer
// Test:  go test ./04-methods/18-exercise-stringer
package main

import "fmt"

type IPAddr [4]byte

// TODO: thêm method "String() string" cho IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
