// Exported names / Tên được export
//
// A name is exported if it begins with a capital letter.
// Một tên được export (nhìn thấy được từ package khác) nếu nó bắt đầu bằng CHỮ HOA.
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi) // Pi được export

	// math.pi would NOT compile: unexported / lowercase.
	// math.pi sẽ KHÔNG biên dịch được: chữ thường nên không export.
	// fmt.Println(math.pi)
}
