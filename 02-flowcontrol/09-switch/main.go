// Switch / Câu lệnh switch
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// KHÁC BIỆT LỚN so với C/Java: Go tự động break sau mỗi case,
	// KHÔNG bị "fall through". Nếu thật sự cần rơi xuống case sau,
	// phải viết tường minh từ khóa fallthrough.
	// case cũng không bắt buộc phải là hằng số nguyên.
}
