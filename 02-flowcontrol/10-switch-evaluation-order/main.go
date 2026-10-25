// Switch evaluation order / Thứ tự tính toán trong switch
package main

import (
	"fmt"
	"time"
)

// switch tính các case theo thứ tự từ trên xuống và DỪNG NGAY khi khớp.
// Các case phía sau sẽ không được tính đến.
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
