// Map literals continued / Khởi tạo map (tiếp)
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// Nếu kiểu ở cấp cao nhất chỉ là tên kiểu, có thể BỎ tên đó trong các phần tử.
// So sánh với bài trước để thấy khác biệt — đây là dạng viết thường gặp.
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
