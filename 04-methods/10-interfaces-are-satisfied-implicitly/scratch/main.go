package main

import (
	"fmt"
	"time"
)

type Speaker interface {
	Speak() string
}

type Mover interface {
	Move() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Dog speak"
}

func (d Dog) Move() string {
	return "Dog move"
}

type Robot struct{}

func (r Robot) Move() string {
	return "Robot move"
}

type Describer interface {
	String() string
}

var _ Speaker = Dog{}

func (r Robot) String() string {
	return "Robot string"
}

func (d Dog) String() string {
	return "Dog string"
}

func main() {
	var speaker Speaker = Dog{}
	var mover Mover = Dog{}
	fmt.Println(speaker.Speak())
	fmt.Println(mover.Move())

	var mover2 Mover = Robot{}
	fmt.Println(mover2.Move())

	var describer Describer = 90 * time.Second
	fmt.Println(describer.String())

	describers := []Describer{
		Robot{},
		Dog{},
	}

	for _, d := range describers {
		fmt.Println(d)
	}

	// var speaker2 Speaker = Robot{}
	// 04-methods\10-interfaces-are-satisfied-implicitly\scratch\main.go:35:25: cannot use Robot{} (value of struct type Robot) as Speaker value in variable declaration: Robot does not implement Speaker (missing method Speak)

	// 04-methods\10-interfaces-are-satisfied-implicitly\scratch\main.go:33:17: cannot use Dog{} (value of struct type Dog) as Speaker value in variable declaration: Dog does not implement Speaker (wrong type for method Speak)
    //             have Speak(bool) string
    //             want Speak() string
	// 04-methods\10-interfaces-are-satisfied-implicitly\scratch\main.go:36:24: cannot use Dog{} (value of struct type Dog) as Speaker value in variable declaration: Dog does not implement Speaker (wrong type for method Speak)
	// 				have Speak(bool) string
	// 				want Speak() string


	// lỗi xảy ra vì hàm Speak có receiver Dog là có thêm đối số truyền vào không phải giống như được định nghĩa trong interface là không có đối số nào cả
}
