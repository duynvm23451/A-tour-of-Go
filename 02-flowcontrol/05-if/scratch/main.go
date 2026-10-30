package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func evenOrOdd(n int) string {
	if n%2 == 0 {
		return "chẵn"
	}
	return "lẻ"
}

func isLeap(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(abs(-5), abs(3))             // 5 3
	fmt.Println(evenOrOdd(4), evenOrOdd(-3)) // chẵn lẻ
	fmt.Println(isLeap(2024), isLeap(1900), isLeap(2000), isLeap(2023))
	// true false true false

	count := 3
	if count > 0 {
		fmt.Println("có phần tử")
	}
}