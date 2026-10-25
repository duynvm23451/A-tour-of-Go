// Short variable declarations / Khai báo biến ngắn gọn
package main

import "fmt"

func main() {
	// Inside a function, := replaces var with implicit type.
	// Bên trong hàm, := thay cho var và tự suy kiểu.
	// LƯU Ý: := chỉ dùng được TRONG hàm, không dùng ở cấp package.
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
