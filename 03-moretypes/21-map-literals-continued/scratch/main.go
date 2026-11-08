package main

import "fmt"

func lookup(scores map[string]float64, name string) {
	if v, ok := scores[name]; ok {
		fmt.Printf("%v : %v\n",name,v )
	} else {
		fmt.Println("Không có trong lớp")
	}
}

type Vertex struct {
	Lat float64
	Long float64
}

func main() {
	scores := map[string]float64{"An": 0, "Binh": 8}
	lookup(scores, "An")
	lookup(scores, "Binh")
	lookup(scores, "Cong")

	delete(scores, "Binh")
	lookup(scores, "Binh")

	delete(scores, "Không có") // không lỗi

	cities := map[string]Vertex{
		"Ha Noi" : {2.4, 3.6},
	}

	city := cities["Ha Noi"]
	city.Lat = 5
	cities["Ha Noi"] = city

	fmt.Println(cities)
}