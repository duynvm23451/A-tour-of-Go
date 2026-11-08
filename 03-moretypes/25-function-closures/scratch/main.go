package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func apply(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = fn(v)
	}
	return result
}

func multiplier(n int) func(int) int {
	return func(value int) int {
		return n * value

	}
}

func main() {
	a, b := counter(), counter()

	fmt.Println(a()) //1
	fmt.Println(a()) //2
	fmt.Println(a()) //3
	fmt.Println(b()) //1

	fmt.Println(a()) //4

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(apply([]int{1, 2, 3}, double))   // mong đợi: [2 4 6]
	fmt.Println(apply([]int{1, 2, 3}, triple))   // mong đợi: [3 6 9]

	total := 0
	add := func(x int) { total += x }
	add(5) // total = 5
	add(10) // total = 15
	fmt.Println(total) // 15

}