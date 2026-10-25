// Switch with no condition / switch không có điều kiện
package main

import (
	"fmt"
	"time"
)

// switch không điều kiện tương đương switch true.
// Đây là cách viết sạch để thay cho chuỗi if - else if - else dài.
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
