// Slice literals / Khởi tạo slice
package main

import "fmt"

// Slice literal giống array literal nhưng KHÔNG ghi độ dài.
// Nó tạo ra mảng nền rồi trả về slice trỏ vào mảng đó.
func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
