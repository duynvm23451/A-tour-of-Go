package main

import "fmt"

func dayName(n int) string {
	switch n {
	case 1:
		return "Chủ nhật"
	case 2:
		return "Thứ hai"
	case 3:
		return "Thứ ba"
	case 4:
		return "Thứ tư"
	case 5:
		return "Thứ năm"
	case 6:
		return "Thứ sáu"
	case 7:
		return "Thứ bảy"
	default:
		return "Không hợp lệ"
	}
}

func isWeekend(n int) bool {
	switch n {
	case 1, 7:
		return true
	default:
		return false
	}
}

func daysIn(month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		return 28
	default:
		return 0
	}
}

// switch n := 2; n {
// case 1:
// 	fmt.Println("một")
// case 2:
// 	fmt.Println("hai")
// 	fallthrough
// case 3:
// 	fmt.Println("ba")
// case 4:
// 	fmt.Println("bốn")
// }
// In ra:
// hai
// ba

func main() {
	fmt.Println(dayName(2), dayName(5))
	fmt.Println(isWeekend(3), isWeekend(7))
	fmt.Println(daysIn(7))
}
