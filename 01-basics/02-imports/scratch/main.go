package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Printf("Dice roll: %d\n", rand.Intn(6)+1)
	fmt.Printf("Square root of 2 is %.3f\n", math.Sqrt(2))
	fmt.Printf("Pi is approximately %.4f\n", math.Pi)
}
