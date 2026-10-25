// Exercise: Equivalent Binary Trees / Bài tập: Hai cây nhị phân tương đương
//
// ĐỀ BÀI
// Nhiều cấu trúc cây khác nhau có thể chứa CÙNG một dãy giá trị.
// Ví dụ, hai cây dưới đây đều chứa 1, 1, 2, 3, 5, 8, 13:
//
//	    3               8
//	   / \             / \
//	  1   8          3    13
//	 / \   \        / \
//	1   2   13     1   5
//	       /      / \
//	      5      1   2
//
// Hàm kiểm tra hai cây có chứa cùng dãy giá trị hay không, trong hầu hết
// ngôn ngữ, là khá rắc rối. Ta sẽ dùng concurrency và channel của Go
// để viết một lời giải đơn giản.
//
// Bước 1: Cài đặt hàm Walk.
//
// Bước 2: Kiểm thử Walk. Hàm tree.New(k) sinh ra cây nhị phân được sắp xếp
// ngẫu nhiên chứa các giá trị k, 2k, 3k, ..., 10k. Tạo channel ch mới rồi
// chạy "go Walk(tree.New(1), ch)", sau đó đọc và in 10 giá trị từ channel.
// Kết quả phải là 1, 2, 3, ..., 10.
//
// Bước 3: Cài đặt hàm Same, dùng Walk để xác định t1 và t2 có cùng giá trị không.
//
// Bước 4: Kiểm thử Same. Same(tree.New(1), tree.New(1)) phải trả về true,
// còn Same(tree.New(1), tree.New(2)) phải trả về false.
//
// Chạy:  go run ./06-concurrency/07-exercise-equivalent-binary-trees
// Test:  go test ./06-concurrency/07-exercise-equivalent-binary-trees
package main

import (
	"fmt"

	"tour/internal/tree"
)

// Walk duyệt cây t theo thứ tự TĂNG DẦN và gửi mọi giá trị vào channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// TODO: duyệt trung tự (in-order): Left, Value, Right.
	// Chú ý: hàm đệ quy không thể tự close ở mỗi lần gọi,
	// nên hãy tách phần đệ quy ra một hàm riêng và chỉ close ở đây.
	close(ch)
}

// Same trả về true nếu t1 và t2 chứa cùng một dãy giá trị.
func Same(t1, t2 *tree.Tree) bool {
	// TODO: chạy hai Walk trong hai goroutine rồi so sánh từng cặp giá trị.
	return false
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
