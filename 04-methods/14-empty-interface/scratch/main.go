package main

import "fmt"

type Book struct {
	Title string
	Pages int
}

func printAll(items []any) {
	for i, v := range items {
		fmt.Printf("[%v] value=%v type=%T\n", i, v, v)
	}
}

func same(a, b any) bool {
	return a == b
}

func main() {
	items := []any{1, "Go", 3.14, true, nil, []int{1, 2, 3}, Book{Title: "Go 101", Pages: 300}}
	printAll(items)

	profile := map[string]any{
		"name": "Duy",
		"age": 25,
		"height": 1.72,
		"isAdmin": true,
		"tags": []string{"student", "class"},
		"address": map[string]any {
			"city": "Okinawa",
			"zip": 23,
		},
	}

	for k, v := range profile {
		fmt.Printf("%v: %v (%T)\n", k, v, v)
	}

	// profile["age"] + 1
	// 04-methods\14-empty-interface\scratch\main.go:36:2: invalid operation: profile["age"] + 1 (mismatched types any and untyped int)
	// profile["age"] có kiểu là any nên không thể cộng thêm 1

	fmt.Println(same(1, 1)) // đúng thì value đều là 1 và type là int
	fmt.Println(same(1, int64(1))) // sai vì valule thì đúng là 1 nhưng một cái type int, một cái type int64
	fmt.Println(same(1, 1.0)) // sai vì một cái type int một cái type float64
	fmt.Println(same("go", "go")) // đúng
	fmt.Println(same(nil, nil)) // đúng

}