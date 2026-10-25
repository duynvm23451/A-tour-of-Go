// Constants / Hằng số
package main

import "fmt"

// Constants are declared like variables, but with the const keyword.
// Hằng số khai báo giống biến nhưng dùng từ khóa const.
// Hằng số có thể là ký tự, chuỗi, bool hoặc số. KHÔNG dùng := với const.
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
