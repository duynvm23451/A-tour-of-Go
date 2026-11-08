package main

import (
	"fmt"
)

func sum(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}

func findEven(nums []int) []int {
	var result []int
	for _, v := range nums {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

type Student struct {
	Name  string
	Age   int
	Score float64
}

func bonusB(students []Student) {
	for i := 0; i < len(students); i++ {
		students[i].Score++
	}
}

func bonusA(students []Student) {
	for _, st := range students {
		st.Score++
	}
}

func printBoard(board [][]string) {
	for _, x := range board {
		for _, y := range x {
			fmt.Print(y, " ")
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println(sum([]int{2,3,4,5,6}))
	fmt.Println(findEven([]int{2,3,4,5,6}))

	class := []Student{
		{Name: "An", Age: 20, Score: 7},
		{Name: "Binh", Age: 22, Score: 8},
		{Name: "Chien", Age: 24, Score: 9},
	}

	board := [][]string{
		{"X", "_", "_"},
		{"_", "X", "_"},
		{"_", "_", "X"},
	}

	bonusB(class) // thay đổi vì dùng chỉ mục
	bonusA(class) // không thay đổi vì dùng range v chỉ sao chép giữ liệu, không thay đổi giá trị gốc
	fmt.Println(class)
	printBoard(board)
}