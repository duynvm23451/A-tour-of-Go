package main

import "fmt"

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	printSlice(s) // len=8 cap=8
	s = s[2:5]
	printSlice(s) // len=3 cap=6
	s = s[:cap(s)]
	printSlice(s) // len=6 cap=6
	s = s[3:]
	printSlice(s) // len=3 cap=3
	s = s[:0]
	printSlice(s) // len=0 cap=3

	// không thể lấy lại phần tử a từ s vì mép trái giờ dịch hết đến f

	t := []int{1, 2, 3, 4, 5}
	u := t[1:3] // 2 3
	fmt.Println(u[:4])
	// fmt.Println(u[:5]) // không chạy được vì giờ cap của u chỉ bằng 4 không thể truy cập vào phần tử thứ 5

	v := u[:4]
	v[3] = 99
	fmt.Println(t)
	// v = 2 3 4 99
	
}