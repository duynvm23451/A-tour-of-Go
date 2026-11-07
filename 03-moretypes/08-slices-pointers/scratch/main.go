package main

import "fmt"

func double(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] *= 2
	}
}

func main() {
	nums := [6]int{10, 20, 30, 40, 50, 60}
	x := nums[1:4]
	y := nums[3:6]
	//    10, 20, 30, 40, 50, 60
	// a      ----------
	// b              ----------
	// chồng phần thử 40
	// phần tử này chỉ mục trong x là 2, trong y là 0
	fmt.Println(x, y) // [20 30 40] [40 50 60]
	x[2] = 999
	fmt.Println(x, y) // [20 30 999] [999 50 60]
	fmt.Println(nums) // [10 20 30 999 50 60]
	double(y)
	fmt.Println(x, y) // [20 30 1998] [1998 100 120]
	fmt.Println(nums) // [10 20 30 1998 100 120]
	// slice là một view của mảng gốc, nên sửa slice sẽ sửa luôn mảng gốc, không cần con trỏ

	arrCopy := nums                       // mảng
	sliceCopy := x                        // slice
	arrCopy[0] = -1                       // sửa arrCopy không sửa nums
	sliceCopy[0] = -2                     // sửa sliceCopy sẽ sửa nums vì sliceCopy là một view của nums
	fmt.Println(arrCopy, sliceCopy, nums) // [-1 20 30 1998 100 120] [-2 30 1998] [-2 30 1998 100 120]

	// câu đó thì phần tử nums[1] = 7 (đã bị thay đổi)
}
