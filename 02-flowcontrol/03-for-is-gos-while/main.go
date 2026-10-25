// For is Go's "while" / for chính là while của Go
package main

import "fmt"

func main() {
	// Bỏ luôn cả hai dấu ; thì for trở thành while.
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
