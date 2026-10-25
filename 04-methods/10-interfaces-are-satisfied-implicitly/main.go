// Interfaces are satisfied implicitly / Interface được thỏa mãn ngầm định
package main

import "fmt"

// A type implements an interface by implementing its methods.
// There is NO "implements" keyword.
//
// Một kiểu cài đặt interface đơn giản bằng cách có đủ method.
// KHÔNG có từ khóa implements, không cần khai báo ý định gì cả.
//
// Nhờ vậy, interface có thể được định nghĩa Ở PHÍA NGƯỜI DÙNG,
// tách rời hoàn toàn khỏi package định nghĩa kiểu.
// Đây là khác biệt lớn nhất giữa interface của Go và của Java/C#.
type I interface {
	M()
}

type T struct {
	S string
}

// T cài đặt I mà không hề "biết" là mình đang cài đặt I.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
