// Packages / Gói
//
// Every Go program is made up of packages. Programs start running in package main.
// Mọi chương trình Go đều gồm các package. Chương trình bắt đầu chạy ở package main.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// By convention, the package name is the last element of the import path:
	// "math/rand" -> package rand.
	// Theo quy ước, tên package là phần cuối của import path.
	fmt.Println("My favorite number is", rand.Intn(10))
}
