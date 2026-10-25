package main

import (
	"fmt"
	"strconv"
)

func chiaLayDu(a, b int) (int, int) {
	return a / b, a % b
}

func minMax(a, b, c int) (int, int) {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return min, max
}

func doiSo(s string) (int, error) {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return result * 2, nil
}

func main() {
	thuong, du := chiaLayDu(17, 5)
	fmt.Printf("17/5 = %d du %d\n", thuong, du)
	min, max := minMax(8, 3, 20)
	fmt.Printf("min= %d max= %d\n", min, max)
	if result, err := doiSo("123"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
