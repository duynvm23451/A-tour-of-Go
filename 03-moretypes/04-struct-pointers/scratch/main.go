package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Student struct {
	Name  string
	Age   int
	Score float64
}

func birthday(s *Student) {
	s.Age++
}

func ex1() {
	v := Vertex{1, 2}
	p := &v // p là con trỏ đến v
	(*p).X = 10
	p.Y = 20       // có thể bỏ dấu * vì Go tự động dereference
	fmt.Println(v) // {10 20}
}

func ex2() {
	student := Student{"An", 20, 7.5}
	birthday(&student) // truyền địa chỉ của student
	fmt.Printf("%+v\n", student)
}

func NewStudent(name string, age int, score float64) *Student {
	return &Student{
		Name:  name,
		Age:   age,
		Score: score,
	}
}

func ex3() {
	student := NewStudent("An", 20, 7.5)
	student.Score = 8
	fmt.Printf("%+v\n", student)
}

func ex4() {
	a := Vertex{1, 2}
	p := &a
	b := a
	p.X = 100
	fmt.Println(a.X, b.X, p.X) // 100 1 100
}

func rename(s *Student, newName string) {
	if s == nil {
		fmt.Println("Không có sinh viên")
		return
	}
	s.Name = newName
}

func ex5() {
	var student *Student
	rename(student, "Bình") // Không có sinh viên
	fmt.Printf("%+v\n", student) // <nil>
	student = NewStudent("An", 20, 7.5)
	rename(student, "Bình")
	fmt.Printf("%+v\n", student) // &{Name:Bình Age:20 Score:7.5}	
}
