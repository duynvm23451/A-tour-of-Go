// Function closures / Bao đóng
package main

import "fmt"

// A closure is a function value that references variables from outside its body.
// Closure là một giá trị hàm tham chiếu tới biến nằm NGOÀI thân hàm nó.
// Hàm đó có thể đọc và gán cho biến được tham chiếu; nói cách khác,
// hàm bị "gắn chặt" (bound) vào các biến đó.
//
// Ở đây mỗi lần gọi adder() sẽ tạo ra một biến sum RIÊNG.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
