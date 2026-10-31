package main

import "fmt"

func main() {
	n := 987654

	count := 0

	for n > 0 {
		n /= 10
		count++
	}

	fmt.Println(count)

	p := 1

	for p*3 <= 500 {
		p *= 3
	}

	fmt.Println(p * 3)

}
