package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func fib(x int) int{
	a, b := 0, 1
	for i := 0 ; i < x; i++ {
		a, b = b, a+b
	}
	return a
}
func main() {
	fmt.Println(gcd(1071, 462))
	fmt.Println(fib(40)) 
}