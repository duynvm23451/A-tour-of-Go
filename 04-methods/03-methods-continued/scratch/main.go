package main

import "fmt"

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	result := (c * 9 / 5) + 32

	return float64(result)
}

func main() {
	fmt.Println(Celsius(100).ToFahrenheit())
}