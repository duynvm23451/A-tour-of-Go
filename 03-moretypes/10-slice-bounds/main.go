// Slice defaults / Giá trị mặc định khi cắt
package main

import "fmt"

// Khi cắt, có thể bỏ cận trên hoặc cận dưới để dùng mặc định:
// cận dưới mặc định là 0, cận trên mặc định là len(slice).
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
