// Variables / Biến
package main

import "fmt"

// "var" declares a list of variables; the type comes last.
// "var" khai báo danh sách biến; kiểu đứng cuối.
// var ở cấp package dùng được trong mọi hàm của package đó.
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
