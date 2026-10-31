package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		if (i % 2) == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	for i := 1; i <= 10; i++ {
		fmt.Println("7 *", i, "=", 7*i)
	}

	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}
}
