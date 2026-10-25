// Maps / Bảng ánh xạ
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// Zero value của map là nil. Map nil KHÔNG có key và KHÔNG thêm key được
// (ghi vào map nil sẽ panic) — khác với slice nil. Phải dùng make trước.
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
