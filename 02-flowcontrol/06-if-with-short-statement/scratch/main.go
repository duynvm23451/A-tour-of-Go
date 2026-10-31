package main

import "fmt"

func bigSquare(n int) bool {
	if sq := n * n; sq > 100 {
		return true
	}
	return false
}

func describe(n int) string {
	if r := n % 3; r == 0 {
		return "chia hết cho 3"
	} else {
		return "không chia hết cho 3"
	}
}

func isLeap(year int) bool {
	if r := year % 400; r == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

func main() {
	fmt.Println(bigSquare(5), bigSquare(11)) // false true
	fmt.Println(describe(9), describe(10))   // chia hết cho 3 / không chia hết cho 3
	fmt.Println(isLeap(2024), isLeap(1900), isLeap(2000), isLeap(2023))
	// true false true false
}
