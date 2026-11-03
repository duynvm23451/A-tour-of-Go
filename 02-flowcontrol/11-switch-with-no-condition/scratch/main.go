package main

import (
	"fmt"
	"time"
)

func classify(i int) string {
	switch {
	case i%15 == 0:
		return "FizzBuzz"
	case i%3 == 0:
		return "Fizz"
	case i%5 == 0:
		return "Buzz"
	default:
		return "số thường"
	}
}

func rank(score float64) string {
	switch {
	case score > 10.0:
		return "Điểm không hợp lệ"
	case score >= 8.5:
		return "Giỏi"
	case score >= 7.0:
		return "Khá"
	case score >= 5.0:
		return "Trung bình"
	case score >= 0.0:
		return "Yếu"
	default:
		return "Điểm không hợp lệ"
	}
}

func main() {
	fmt.Println(classify(15)) // Output: FizzBuzz
	fmt.Println(rank(9.0))    // Output: Giỏi
	fmt.Println(classify(9))  // Output: Fizz
	fmt.Println(classify(10)) // Output: Buzz
	fmt.Println(classify(7))  // Output: số thường
	fmt.Println(rank(11.0))   // Output: Điểm không hợp lệ

	switch h := time.Now().Hour(); {
	case h < 5:
		fmt.Println("Khuya rồi, đi ngủ đi!")
	case h < 12:
		fmt.Println("Good morning!")
	case h < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
