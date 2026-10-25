// For / Vòng lặp for
package main

import "fmt"

// Go has only ONE looping construct: the for loop.
// Go chỉ có DUY NHẤT một cấu trúc lặp: for. Không có while, không có do-while.
func main() {
	sum := 0
	// Ba thành phần cách nhau bởi dấu ;   init; condition; post
	// KHÔNG có dấu ngoặc tròn bao quanh, nhưng dấu { } thì BẮT BUỘC.
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
