// Basic types / Các kiểu cơ bản
package main

import (
	"fmt"
	"math/cmplx"
)

// Các kiểu cơ bản của Go:
//
//	bool
//	string
//	int  int8  int16  int32  int64
//	uint uint8 uint16 uint32 uint64 uintptr
//	byte // alias của uint8
//	rune // alias của int32, đại diện một Unicode code point
//	float32 float64
//	complex64 complex128
//
// int, uint, uintptr rộng 32 bit trên hệ 32-bit và 64 bit trên hệ 64-bit.
// Khi cần số nguyên, hãy dùng int trừ khi có lý do cụ thể để chọn kiểu khác.
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
