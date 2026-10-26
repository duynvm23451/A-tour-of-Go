package main

import "fmt"

func main() {
	// Declare a variable named "i" of type int.
	// Khai báo một biến tên là "i" kiểu int.
	varBool := true
	varString := "Hello"
	varInt := 42
	varInt8 := int8(127)
	varUInt8 := uint8(255)
	varFloat := 3.14
	varRune := 'A'
	varByte := byte(255)

	fmt.Printf("varBool: %v, type: %T\n", varBool, varBool)
	fmt.Printf("varString: %v, type: %T\n", varString, varString)
	fmt.Printf("varInt: %v, type: %T\n", varInt, varInt)
	fmt.Printf("varInt8: %v, type: %T\n", varInt8, varInt8)
	fmt.Printf("varUInt8: %v, type: %T\n", varUInt8, varUInt8)
	fmt.Printf("varFloat: %v, type: %T\n", varFloat, varFloat)
	fmt.Printf("varRune: %v, type: %T\n", varRune, varRune)
	fmt.Printf("varByte: %v, type: %T\n", varByte, varByte)

	varInt8++
	fmt.Printf("varInt8: %v, type: %T\n", varInt8, varInt8)

	myName := "Duy"
	myFullName := "Mạnh Duy"
	myCity := "Hà Nội"
	fmt.Println("len(myName):", len(myName))
	fmt.Println("len(myFullName):", len(myFullName))
	fmt.Println("len(myCity):", len(myCity))

	fmt.Println("len([]rune(myName)):", len([]rune(myName)))
	fmt.Println("len([]rune(myFullName)):", len([]rune(myFullName)))
	fmt.Println("len([]rune(myCity)):", len([]rune(myCity)))

	fmt.Printf("%v %c %T\n", 'A', 'A', 'A')
}