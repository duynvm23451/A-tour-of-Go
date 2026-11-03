package main

import "fmt"

func work() {
	fmt.Println("mở cửa")
	defer fmt.Println("đóng cửa")
	fmt.Println("mở đèn")
	defer fmt.Println("tắt đèn")
	fmt.Println("mở máy tính")
	defer fmt.Println("tắt máy tính")
}

func handle(i int) {
	fmt.Println("mở file", i)
	defer fmt.Println("xử lý xong", i)
}

func main() {
	defer fmt.Println("A")
	fmt.Println("B")
	defer fmt.Println("C")
	fmt.Println("D")
	defer fmt.Println("E")

	// Output:
	// B
	// D
	// E
	// C
	// A

	work()

	for i := 1; i <= 3; i++ {
		handle(i)
	}
	fmt.Println("xong")

}
