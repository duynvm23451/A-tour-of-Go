package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func MoveFunc(p *Point, dx, dy int) {
	p.X += dx
	p.Y += dy
}

func Distance(p Point) int {
	result := math.Abs(float64(p.X)) + math.Abs(float64(p.Y))
	return int(result)
}

func (p Point) DistanceRewrite() int {
	result := math.Abs(float64(p.X)) + math.Abs(float64(p.Y))
	return int(result)
}


func main() {
	a := Point{0, 0}

	b := &Point{0, 0}

	a.Move(1, 2)

	b.Move(3, 4)

	fmt.Println(a) // {X: 1, Y: 2} // go tự thêm &
	fmt.Println(b) // {X: 3, Y: 4}

	// MoveFunc(a, 1, 1)
	// 04-methods\05-methods-pointers-explained\scratch\main.go:31:11: cannot use a (variable of struct type Point) as *Point value in argument to MoveFunc

	MoveFunc(&a, 1, 1)
	MoveFunc(b, 1, 1)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(Distance(a))

	fmt.Println(Distance(*b))

	fmt.Println(b.DistanceRewrite())
}