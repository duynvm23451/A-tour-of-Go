// Arrays / Mảng
package main

import "fmt"

// The type [n]T is an array of n values of type T.
// Kiểu [n]T là mảng gồm n phần tử kiểu T.
//
// QUAN TRỌNG: độ dài là MỘT PHẦN CỦA KIỂU. [4]int và [5]int là hai kiểu khác nhau.
// Vì thế mảng không thể thay đổi kích thước. Trong thực tế, Go dùng slice nhiều hơn.
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Mảng là GIÁ TRỊ: gán mảng sẽ copy toàn bộ phần tử.
	b := primes
	b[0] = 100
	fmt.Println(primes, b) // primes không đổi
}
