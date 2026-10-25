// Lời giải: Equivalent Binary Trees
package main

import (
	"fmt"

	"tour/internal/tree"
)

// walk là phần đệ quy: duyệt trung tự (in-order) nên giá trị đi ra
// theo thứ tự TĂNG DẦN. Nó KHÔNG close channel, vì mỗi lần gọi đệ quy
// mà close thì channel sẽ bị đóng sớm.
func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

// Walk là lớp vỏ bên ngoài: chạy đệ quy xong thì đóng channel
// để bên nhận biết là đã hết dữ liệu.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

// Same so sánh hai cây bằng cách duyệt song song và đối chiếu từng cặp.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// Hai cây phải cạn dữ liệu cùng lúc.
		if ok1 != ok2 {
			return false
		}
		// Cả hai đã đóng -> duyệt hết mà không thấy khác biệt.
		if !ok1 {
			return true
		}
		if v1 != v2 {
			return false
		}
	}

	// Ghi chú: nếu trả về false sớm, goroutine Walk còn lại sẽ bị kẹt
	// mãi mãi ở lệnh gửi (goroutine leak). Trong code thật hãy dùng
	// channel quit hoặc context để báo cho chúng dừng lại.
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println("Same(tree.New(1), tree.New(1)) =", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2)) =", Same(tree.New(1), tree.New(2)))
}
