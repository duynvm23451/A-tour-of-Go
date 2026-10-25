// Slices / Lát cắt
package main

import "fmt"

// An array has a fixed size. A slice is dynamically-sized and far more common.
// Mảng có kích thước cố định. Slice thì linh hoạt và được dùng nhiều hơn hẳn.
//
// Kiểu []T là một slice chứa các phần tử kiểu T.
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Cú pháp cắt: a[low : high] — bao gồm low, KHÔNG bao gồm high.
	var s []int = primes[1:4]
	fmt.Println(s) // [3 5 7]
}
