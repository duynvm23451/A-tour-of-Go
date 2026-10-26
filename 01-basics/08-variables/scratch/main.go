package main

import "fmt"

var (
	ten         string
	tuoi        int
	diem        float64
	daTotNghiep bool
)

var soLanGoi int

func main() {
	// Declare a variable named "i" of type int.
	// Khai báo một biến tên là "i" kiểu int.
	fmt.Printf("%q %v %v %v\n", ten, tuoi, diem, daTotNghiep)
	tang()
	tang()
	tang()
	fmt.Println("Số lần gọi hàm tang():", soLanGoi)
}

func tang() {
	soLanGoi++
}
