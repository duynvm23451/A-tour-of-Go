package main

import (
	"fmt"
	"slices"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

func average(students []Student) float64 {
	if len(students) == 0 {
		return 0
	}
	sum := 0.0
	for i := 0; i < len(students); i++ {
		sum += students[i].Score
	}
	return sum / float64(len(students))
}

func bonus(students []Student) {
	for i := 0; i < len(students); i++ {
		students[i].Score++
	}
}

func main() {
	a := [...]string{"go", "is", "fun"}
	b := []string{"go", "is", "fun"}
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	c := a
	d := b
	c[0] = "GO"
	d[0] = "GO"
	fmt.Println(a) // [go is fun]
	fmt.Println(b) // [GO is fun]

	class := []Student{
		{Name: "An", Age: 20, Score: 7},
		{Name: "Binh", Age: 22, Score: 8},
		{Name: "Chien", Age: 24, Score: 9},
	}
	fmt.Printf("%+v\n", class)

	fmt.Println(average(class))

	bonus(class)
	fmt.Println(class) // Điểm có tăng vì slice không cần con trỏ để sửa giá trị

	// []int{1, 2} == []int{1, 2} // sai vì chỉ array mới so sánh được kiểu này, slice chỉ có thể compare với nil
	fmt.Println(slices.Equal([]int{1, 2}, []int{1, 2}))
}