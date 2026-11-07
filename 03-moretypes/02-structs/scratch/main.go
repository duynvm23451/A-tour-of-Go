package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

func main() {
	student := Student{
		Name:  "Alice",
		Age:   20,
		Score: 85.5,
	}

	var studentNil Student
	fmt.Printf("%+v\n", student)
	fmt.Printf("%+v\n", studentNil)

	rectangle := Rectangle{
		Width:  5.0,
		Height: 3.0,
	}
	fmt.Printf("Area of rectangle: %.2f\n", rectangle.area())

	// a := Vertex{1, 2}
	// b := Vertex{1, 2}
	// c := Vertex{2, 1}
	// fmt.Println(a == b, a == c) // true false

	// st2 := Student{
	// 	Name:  "Bob"
	// }
	// => error
}
