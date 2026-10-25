// Generic types / Kiểu generic
package main

import "fmt"

// Kiểu cũng có thể nhận tham số kiểu. Đây là danh sách liên kết đơn
// chứa giá trị thuộc kiểu T bất kỳ.
type List[T any] struct {
	next *List[T]
	val  T
}

// Method của kiểu generic khai báo lại tham số kiểu ở receiver.
func (l *List[T]) Push(v T) *List[T] {
	return &List[T]{next: l, val: v}
}

func (l *List[T]) All() []T {
	var out []T
	for node := l; node != nil; node = node.next {
		out = append(out, node.val)
	}
	return out
}

func main() {
	var ints *List[int]
	ints = ints.Push(1).Push(2).Push(3)
	fmt.Println(ints.All())

	var strs *List[string]
	strs = strs.Push("a").Push("b")
	fmt.Println(strs.All())
}
