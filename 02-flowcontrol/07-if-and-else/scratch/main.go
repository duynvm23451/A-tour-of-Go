package main

import "fmt"

func sign(n int) string {
	if n < 0 {
		return "âm"
	} else if n == 0 {
		return "không"
	} else {
		return "dương"
	}
}

func checkAge(birthYear int) string {
	if age := 2026 - birthYear; age >= 18 {
		return "đủ tuổi"
	} else {
		return fmt.Sprintf("còn thiếu %d năm", 18-age)
	}
}

func grade(score int) string {
	s := score / 10
	if s >= 9 {
		return "A"
	}
	if s >= 7 {
		return "B"
	}
	if s >= 5 {
		return "C"
	}
	return "F"
}

// func say(s string) string {
// 	fmt.Println("đang gọi", s)
// 	return s
// }
// func main() {
// 	fmt.Println(say("A"), say("B"))
// đang gọi A
// đang gọi B
// A
// B
// }

func main() {
	fmt.Println(sign(-5), sign(0), sign(7))
	fmt.Println(checkAge(2008), checkAge(2010))
	fmt.Println(grade(95), grade(72), grade(55), grade(30))
}
