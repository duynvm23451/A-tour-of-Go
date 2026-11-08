package main

import "fmt"

func squares(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i * i
	}
	return arr
}

func squares2(n int) []int {
	arr := make([]int, n)
	for i := range n {
		arr[i] = i*i
	}
	return arr
}

func countdown(n int) {
	for range n {
		fmt.Println("tick")
	}
}

func main() {
	// for i, v := range []int{1, 2, 3} {
	// 	fmt.Println(v)
	// }
	// 03-moretypes\17-range-continued\scratch\main.go:6:6: declared and not used: i
	fmt.Println(squares(5))
	fmt.Println(squares2(5))
	
	countdown(3)
}