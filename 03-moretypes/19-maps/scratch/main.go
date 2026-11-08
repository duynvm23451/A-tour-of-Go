package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score float64
}

func classMap(students []Student) map[string]float64 {
	result := make(map[string]float64)
	for _, v := range students {
		result[v.Name] = v.Score
	}
	return result
}

func main() {
	scores := make(map[string]float64)
	scores["An"] = 9.9
	scores["Binh"] = 8
	scores["Cong"] = 7
	scores["An"] = 5

	fmt.Println(scores)
	fmt.Println(len(scores))
	fmt.Println(scores["Không có"])

	class := []Student{
		{Name: "An", Age: 20, Score: 7},
		{Name: "Binh", Age: 22, Score: 8},
		{Name: "Chien", Age: 24, Score: 9},
	}

	fmt.Println(classMap(class))

	var m map[string]int
	fmt.Println(m["a"], len(m), m == nil)
	// m["a"] = 1
	// panic: assignment to entry in nil map

	// goroutine 1 [running]:
	// main.main()
	// 		D:/go lang/03-moretypes/19-maps/scratch/main.go:40 +0x3d5
	// exit status 2
	// không thể thêm phần tử vào nil map trong khi append với nil slice thì làm được
}