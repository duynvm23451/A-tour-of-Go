// Range continued / range (tiếp)
package main

import "fmt"

func main() {
	pow := make([]int, 10)

	// Bỏ giá trị bằng cách gán vào _ (blank identifier).
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	// Bỏ chỉ số bằng cách viết  for _, value := range ...
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// LƯU Ý: value là BẢN SAO. Sửa value không ảnh hưởng tới slice gốc;
	// muốn sửa thì phải ghi qua chỉ số: pow[i] = ...
}
