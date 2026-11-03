package main

import "fmt"

func ex1() {
	a := 10
	p := &a
	q := p
	*q = 20
	fmt.Println(a, *p, *q)
	// Output:
	// 20 20 20
	fmt.Println(p == q)
	// Output:
	// true
}

func increment(p *int) {
	*p++
}

func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}

func resetValue(x int) {
	x = 0
}

func resetPointer(p *int) {
	*p = 0
}

func main() {
	ex1()
	var a int

	increment(&a)
	increment(&a)
	increment(&a)
	fmt.Println(a) // Output: 3
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y) // Output: 2 1
	var p *int
	fmt.Println(p) // Output: <nil>
	// fmt.Println(*p) // Output: error
	if p != nil {
		fmt.Println(*p)
	}

	resetValue(a)
	fmt.Println(a) // Output: 10
// không thay đổi vì pass by value
	resetPointer(&a)
	fmt.Println(a) // Output: 0
// thay đổi vì pass a pointer
}