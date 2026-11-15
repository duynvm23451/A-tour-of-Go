package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	W, H float64
}

type Circle struct {
	R float64
}

func (r Rectangle) Area() float64 {
	return r.W * r.H
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

func totalArea(shapes []Shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}

type Square struct {
	S float64
}

func (sq *Square) Area() float64 {
	return sq.S * sq.S
}

func main() {
	r := Rectangle{W: 20, H: 30}
	c := Circle{R: 4}
	printArea(r)
	printArea(c)

	c2 := Circle{R: 6}

	shapes := []Shape{r, c, c2}
	fmt.Println(totalArea(shapes))

	shapes = append(shapes, &Square{2})
	// 04-methods\09-interfaces\scratch\main.go:59:26: cannot use Square{…} (value of struct type Square) as Shape value in argument to append: Square does not implement Shape (method Area has pointer receiver)
}