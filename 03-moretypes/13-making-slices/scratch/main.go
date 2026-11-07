package main

import "fmt"

func squares(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i * i
	}
	return arr
}

func main() {
	x := make([]string, 3)  // len=3, cap=3
	y := make([]bool, 2, 6) // len=2, cap=6
	z := y[1:4]             // len=3, cap=5

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// squares(0) không trả về nil mà trả về [] với len=cap=3

	// make([]int, 5, 3) // lỗi vì len > cap

}
