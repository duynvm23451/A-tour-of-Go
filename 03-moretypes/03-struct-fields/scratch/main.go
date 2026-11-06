package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

func birthday(s Student) {
	s.Age++ 
}

func birthdayReturnObj(s Student) Student {
	s.Age++
	return s
}

func main() {
	student := Student{"An", 20, 7.5}
	student.Age++ // truy cập bằng dấu chấm
	student.Score += 1.5
	fmt.Printf("%+v\n", student) // in ra tất cả các trường của struct

	// a := Student{"An", 20, 7.5}
	// b := a
	// b.Name = "Bình"
	// b.Score = 9
	// fmt.Println(a.Name, a.Score) // An 7.5
	// fmt.Println(b.Name, b.Score) // Bình 9

	birthday(student)
	fmt.Printf("%+v\n", student) // Khong thay doi vi birthday nhận vào một bản sao của student

	student = birthdayReturnObj(student)
	fmt.Printf("%+v\n", student)

	// p := &a           // a là Student ở bài 2
	// b := *p           // b nhận gì?
	// b.Name = "Cường"
	// fmt.Println(a.Name) // An


}