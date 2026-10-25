// For continued / for (tiếp)
package main

import "fmt"

func main() {
	// The init and post statements are optional.
	// Câu lệnh init và post là tùy chọn — có thể bỏ đi:
	//
	//	for ; sum < 1000; {
	//		sum += sum
	//	}
	//
	// Lưu ý: gofmt sẽ tự rút gọn dạng trên thành dạng dưới đây,
	// nên trong code thật bạn sẽ luôn thấy dạng ngắn.
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
