package main

import "fmt"

func main() {
	n := 7
	for {
		if n%7 == 0 && n%13 == 0 {
			break
		}
		n++
	}
	fmt.Println("Found it!", n)

	for i := 1; i < 20; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	n = 27
	step := 0
	for {
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		step++
	}

	fmt.Println(step)
}
