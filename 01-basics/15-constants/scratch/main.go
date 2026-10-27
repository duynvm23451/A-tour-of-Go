package main

import "fmt"

const (
	SecondsPerMinute int = 60
	SecondsPerHour   int = 60 * SecondsPerMinute
	SecondsPerDay    int = 24 * SecondsPerHour
)

func main() {
	// Constants / Hằng số

	const limit = 100
	const Max = limit * 2

	fmt.Println("Seconds per minute:", SecondsPerMinute)
	fmt.Println("Seconds per hour:", SecondsPerHour)
	fmt.Println("Seconds per day:", SecondsPerDay)
}