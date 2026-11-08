package main

import (
	"fmt"
	"strings"
)

func findEven(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			result = append(result, nums[i])
		}
	}
	return result
}

func main() {
	fmt.Println(findEven([]int{2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(findEven([]int{2, 3, 4, 5, 6, 7, 8}) == nil)
	x := make([]int, 3, 10)
	y := append(x, 4)
	y[0] = 99
	fmt.Println(x, y) // x: 99, 0, 0 ; y: 99, 0, 0, 4
	
	a := []string{"go", "is"}
	b := []string{"fun", "!"}

	joinedSlice := append(a, b...)
	fmt.Println(strings.Join(joinedSlice, " "))

}