package main

import "fmt"

func ex1() {
	x := 1
	defer fmt.Println("A:", x)
	defer func() { fmt.Println("B:", x) }()
	x = 2
	defer fmt.Println("C:", x)
	x = 3

	// Output:
	// C: 2
	// B: 3
	// A: 1
}

func ex2() {
	x := 10
	defer fmt.Println("x * 2 =", x*2)
	x = 100
	// Output:
	// x * 2 = 20
}

func main() {
	ex1()
	ex2()
}
