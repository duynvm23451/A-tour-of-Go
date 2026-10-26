package main

import (
	"fmt"
	"strconv"
)

func chiaLayDu(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	ten, tuoi, diem, daTotNghiep := "Duy", 25, 8.5, true
	fmt.Printf("%q %T\n", ten, ten)
	fmt.Printf("%v %T\n", tuoi, tuoi)
	fmt.Printf("%v %T\n", diem, diem)
	fmt.Printf("%v %T\n", daTotNghiep, daTotNghiep)

	thuong, du := chiaLayDu(17, 5)
	fmt.Printf("17/5 = %d du %d\n", thuong, du)

	gia := 100
	gia = 99.0

	fmt.Printf("%v %T\n", gia, gia)

	value, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Converted value: %d %T\n", value, value)
	}

	value2, err := strconv.Atoi("abc")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Converted value: %d %T\n", value2, value2)
	}
}
