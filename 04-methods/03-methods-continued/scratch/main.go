package main

import "fmt"

type Celsius float64

type Scores []float64

func (celsius Celsius) ToFahrenheit() float64 {
	result := (celsius * 9 / 5) + 32

	return float64(result)
}

func (scores Scores) Average() float64 {
	sum := 0.0
	for _, v := range scores {
		sum += v
	}
	result := sum / float64(len(scores))
	return result
}

func (scores Scores) Max() float64 {
	max := 0.0
	for _, v := range scores {
		if v >= max {
			max = v
		}
	}
	return max
}

// func (i int) Double() int { return i * 2 }
// 04-methods\03-methods-continued\scratch\main.go:34:9: cannot define new methods on non-local type int

type MyInt int

func (i MyInt) Double() MyInt {
	return i*2
} 
func main() {
	fmt.Println(Celsius(100).ToFahrenheit())

	scores := Scores{7, 8, 9}
	fmt.Println(scores.Average())
	fmt.Println(scores.Max())

	fmt.Println(MyInt(21).Double())
}