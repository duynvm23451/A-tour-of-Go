package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, b := 7, 2

	// a / b = 3, vì cả a và b đều là int
	fmt.Println((float64(a) + float64(b)) / 2) // 4.5

	big := 1000
	fmt.Println(int8(big))

	num, _ := strconv.Atoi("123")
	fmt.Println(num + 1) // 124
}
