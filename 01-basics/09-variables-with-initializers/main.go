// Variables with initializers / Biến kèm giá trị khởi tạo
package main

import "fmt"

// If an initializer is present, the type can be omitted — it is inferred.
// Nếu có giá trị khởi tạo thì có thể bỏ kiểu — Go tự suy ra (type inference).
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
