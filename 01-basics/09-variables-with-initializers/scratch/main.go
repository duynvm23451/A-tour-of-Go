package main

import "fmt"

func chiaLayDu(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	// Declare a variable named "i" of type int.
	// Khai báo một biến tên là "i" kiểu int.
	var ten, tuoi, diem, daTotNghiep = "Duy", 25, 8.5, true
	fmt.Printf("%q %T\n", ten, ten)
	fmt.Printf("%v %T\n", tuoi, tuoi)
	fmt.Printf("%v %T\n", diem, diem)
	fmt.Printf("%v %T\n", daTotNghiep, daTotNghiep)

	var thuong, du = chiaLayDu(17, 5)
	fmt.Printf("17/5 = %d du %d\n", thuong, du)

	var gia = 100
	gia = 99.0

	fmt.Printf("%v %T\n", gia, gia)
}
