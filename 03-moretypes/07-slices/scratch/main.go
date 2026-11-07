package main

import "fmt"

func sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	letters := [7]string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(letters[0:3]) // a b c
	fmt.Println(letters[2:5]) // c d e
	fmt.Println(letters[3:4]) // d
	fmt.Println(letters[6:7]) // g
	fmt.Println(letters[3:3]) // []

	fmt.Println(letters[2:6])

	fmt.Printf("%T\n", letters)
	fmt.Printf("%T\n", letters[2:6])
	// letters là một mảng có kiểu [7]string
	// letters[2:6] là một slice có kiểu []string

	arr := [4]int{1, 2, 3, 4}
	fmt.Println(sum(arr[0:4]))
	fmt.Println(sum(arr[1:3]))
	// khong bi loi

	s := letters[1:4]
	s[0] = "x"
	fmt.Println(letters) // [a x c d e f g]
	// slice sẽ sửa luôn mảng gốc, vì slice là một view của mảng gốc

	// fmt.Println(letters[5:2]) // không được vì low > high
	// fmt.Println(letters[0:8]) // không được vì high > len(letters)
}
