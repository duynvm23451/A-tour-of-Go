// Range / Duyệt bằng range
package main

import "fmt"

// The range form of the for loop iterates over a slice or map.
// Dạng range của for dùng để duyệt slice hoặc map.
// Với slice, mỗi vòng trả về hai giá trị: CHỈ SỐ và BẢN SAO của phần tử.
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
