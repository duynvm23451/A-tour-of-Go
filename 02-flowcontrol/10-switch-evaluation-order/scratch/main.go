package main

import (
	"fmt"
	"time"
)

func check(name string, v int) int {
	fmt.Println("đang kiểm tra", name)
	return v
}

func classify(i int) string {
	switch 0 {
	case i % 15:
		return "FizzBuzz"
	case i % 3:
		return "Fizz"
	case i % 5:
		return "Buzz"
	default:
		return "số thường"
	}
}

func main() {
	switch 5 {
	case check("A", 1):
		fmt.Println("khớp A")
	case check("B", 2):
		fmt.Println("khớp B")
	case check("C", 3):
		fmt.Println("khớp C")
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println("days left to Saturday: ", int(time.Saturday-today))

}
