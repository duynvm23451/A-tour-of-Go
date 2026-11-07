package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

type Vertex struct {
	X int
	Y int
}

type Rectangle struct {
	Height int
	Width  int
}

func main() {
	student1 := Student{"An", 20, 7.5}
	student2 := Student{Age: 21, Name: "Binh", Score: 8.0}
	var student3 Student
	var student4 *Student
	fmt.Printf("%+v\n", student1)
	fmt.Printf("%+v\n", student2)
	fmt.Printf("%+v\n", student3)
	fmt.Printf("%+v\n", student4)
	student5 := Student{Score: 9.0, Name: "Anh", Age: 20}
	student6 := Student{"Anh", 20, 9.0}
	fmt.Println(student5 == student6) // true

	// Vertex{1, Y: 2} // sai vì vừa kiểu khai báo có giá trị vừa kiểu khai báo có tên trường
	// Vertex{1} // sai vì đang kiểu khai báo giá trị phải có đủ và đúng thứ tự
	// Vertex{X: 1, X: 2} // sai vì trùng tên trường
	// Vertex{Z: 1} // sai vì Z không phải là trường của struct Vertex

	r1 := Rectangle{3, 4}
	r2 := Rectangle{Width: 3, Height: 4}
	fmt.Printf("%+v\n", r1)
	fmt.Printf("%+v\n", r2)
	// không có lỗi biên dịch nhưng r1 và r2 không bằng nhau vì thứ tự trường khác nhau

	item := struct {
		Name  string
		Price float64
	}{
		Name:  "Iphone",
		Price: 1000,
	}
	fmt.Printf("%+v\n", item)
}
