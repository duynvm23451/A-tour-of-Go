package main

import "fmt"

func apply(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = fn(v)
	}
	return result
}

func filter(nums []int, keep func(int) bool) []int {
	var result []int

	for _, v := range nums {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {

	arr := []int{2, 3, 4, 5, 6}

	double := func(value int) int {
		return value * 2
	}

	square := func(value int) int {
		return value * value
	}

	fmt.Println(apply(arr, double))
	fmt.Println(apply(arr, square))

	evenArr := filter(arr, func(n int) bool { return n%2 == 0 })  // lọc số chẵn
	moreThan3Arr := filter(arr, func(n int) bool { return n > 3 })    // lọc số lớn hơn 3

	fmt.Println(evenArr)
	fmt.Println(moreThan3Arr)

	ops := map[string]func(float64, float64) float64{
		"+": func (x, y float64) float64  {
			return x + y
		},
		"-": func (x, y float64) float64  {
			return x - y
		},
		"*": func (x, y float64) float64  {
			return x * y
		},
		"/": func (x, y float64) float64  {
			if y == 0 {
				return 0
			}
			return x/ y
		},
	}

	fmt.Println(ops["*"](6, 7))

	if op, ok := ops["%"]; ok {
		fmt.Println(op(6, 7))
	} else {
		fmt.Println("không hỗ trợ phép %")
	}
}