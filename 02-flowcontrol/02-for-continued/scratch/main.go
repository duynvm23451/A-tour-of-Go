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

	i := 1
	for ; i*i < 200; i++ {
	}
	fmt.Println(i)

	n = 27

	step := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		step++
	}
	fmt.Println(step)
}
