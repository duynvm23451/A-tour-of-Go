package main

import "fmt"

func double(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
}

func workDouble(arr *[3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
}

func sum(arr [3]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	var scores [5]float64
	fmt.Println(scores) // [0 0 0 0 0]

	days := [...]string{0: "Monday", 1: "Tuesday", 2: "Wednesday", 3: "Thursday", 4: "Friday", 5: "Saturday", 6: "Sunday"}
	fmt.Println(len(days))

	scores2 := [3]int{1, 2, 3}
	double(scores2)
	fmt.Println(scores2) // [1 2 3]

	scores3 := [3]int{1, 2, 3}
	workDouble(&scores3)
	fmt.Println(scores3) // [2 4 6]

	arr := [4]int{1, 2, 3, 4}
	// fmt.Println(sum(arr)) // error: cannot use arr (variable of type [4]int) as [3]int value in argument to sum

	fmt.Println(arr)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	// fmt.Println(primes[6]) // error
}
