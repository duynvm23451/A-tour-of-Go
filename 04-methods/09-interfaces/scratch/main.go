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

func main() {
	r := Rectangle{W: 20, H: 30}
	c := Circle{R: 4}
	printArea(r)
	printArea(c)
}