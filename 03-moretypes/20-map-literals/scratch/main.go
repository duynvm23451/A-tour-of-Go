package main

import "fmt"

type Vertex struct {
	Lat float64
	Long float64
}

func main() {
	scores := map[string]float64{
		"An":   9.9,
		"Binh": 8,
		"Cong": 7,
	}
	scores["Duy"] = 9

	fmt.Println(scores, len(scores))

	cities := map[string]Vertex{
		"Ha Noi" : {2.4, 3.6},
	}

	tests := map[string]int {
		"A": 2,
		// "A": 3,
	}
	// 03-moretypes\20-map-literals\scratch\main.go:26:3: duplicate key "A" in map literal

	fmt.Println(tests)

	fmt.Println(cities["Ha Noi"].Lat)

	subjects := map[string][]string{
		"An":   {"Toán", "Lý"},
		"Binh": {"Văn"},
	}

	fmt.Println(subjects["An"][1])
}