// Type parameters / Tham số kiểu
package main

import "fmt"

// Go functions can be written to work on multiple types using type parameters.
// Hàm Go có thể làm việc với nhiều kiểu nhờ tham số kiểu (từ Go 1.18).
//
// Ở đây S là tham số kiểu, bị ràng buộc bởi "comparable" — nghĩa là
// chỉ chấp nhận kiểu dùng được toán tử == và !=.
func Index[S comparable](s []S, x S) int {
	for i, v := range s {
		// v và x có kiểu S, tức là kiểu comparable, nên dùng == được.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{10, 20, 15, 3}
	fmt.Println(Index(si, 15)) // kiểu được suy ra, không cần Index[int]

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	// Ghi chú: Go 1.21+ đã có sẵn slices.Index trong stdlib.
	// Trong code thật, hãy dùng package slices và maps trước khi tự viết generic.
}
