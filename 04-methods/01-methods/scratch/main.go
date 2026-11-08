package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Student struct {
	Name string
	Age int
	Score float64
}

func (s Student) isPassed() bool {
	return s.Score >= 5
}

func (s Student) Bonus() {
	s.Score++
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())

	class := []Student{
		{Name: "An", Age: 20, Score: 9},
		{Name: "Binh", Age: 20, Score: 8},
		{Name: "Cong", Age: 20, Score: 3},
	}

	for _, v := range class {
		if v.isPassed() {
			fmt.Println(v.Name)
		}
	}

	class[0].Bonus()
	fmt.Println(class[0]) // không thay đổi vì receiver chỉ là bản sao
}