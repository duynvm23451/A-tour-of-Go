// Bonus: iota
//
// Bài này không có trong Tour of Go gốc, nhưng iota xuất hiện khắp nơi
// trong code Go thực tế nên nắm sớm sẽ rất tiện.
package main

import "fmt"

// iota là bộ đếm, reset về 0 ở đầu mỗi khối const và tăng 1 sau mỗi dòng.
type Weekday int

const (
	Sunday    Weekday = iota // 0
	Monday                   // 1
	Tuesday                  // 2
	Wednesday                // 3
	Thursday                 // 4
	Friday                   // 5
	Saturday                 // 6
)

// Mẫu rất phổ biến: dùng iota để sinh các đơn vị theo lũy thừa của 2.
const (
	_  = iota             // bỏ qua giá trị 0
	KB = 1 << (10 * iota) // 1 << 10
	MB                    // 1 << 20
	GB                    // 1 << 30
)

func main() {
	fmt.Println(Sunday, Wednesday, Saturday)
	fmt.Println(KB, MB, GB)
}
