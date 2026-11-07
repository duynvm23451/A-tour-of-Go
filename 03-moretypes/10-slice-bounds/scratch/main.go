package main

import "fmt"

func firstN(s []int, n int) []int {
	return s[:n]
}

func lastN(s []int, n int) []int {
	return s[len(s)-n:]
}

func split(s []int, n int) ([]int, []int) {
	return s[:n], s[n:]
}

func sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	nums := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(nums[:3])
	fmt.Println(nums[1:])
	fmt.Println(nums[:len(nums)-1])
	fmt.Println(nums[len(nums)-2:])
	fmt.Println(nums[:])

	// s := []string{"a", "b", "c", "d", "e", "f", "g"}
	// s = s[2:] // c d e f g
	// s = s[:4] // c d e f
	// s = s[1:3] // d e
	// s = s[1:] // e

	arr := [4]int{1, 2, 3, 4}
	fmt.Println(sum(arr[:]))

	// fmt.Println(sum(arr)) // error because sum accept a slice not a arr

	// call firstN(nums, 100) will be error because it out of range slice
}
