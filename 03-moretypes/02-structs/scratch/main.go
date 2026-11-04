package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

func main() {
	student := Student{
		Name:  "Alice",
		Age:   20,
		Score: 85.5,
	}
	fmt.Printf("%+v\n", student)
}