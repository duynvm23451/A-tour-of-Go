package main

import (
	"fmt"
)

func chuVi(dai, rong float64) float64 {
	return (dai + rong) * 2
}

func dienTich(dai, rong float64) float64 {
	return dai * rong
}

func apDungHai(f func(float64, float64) float64, a, b float64) float64 {
	return f(a, b)
}

func main() {
	fmt.Println("Chu vi:", apDungHai(chuVi, 10, 5))
	fmt.Println("Diện tích:", apDungHai(dienTich, 10, 5))
}
