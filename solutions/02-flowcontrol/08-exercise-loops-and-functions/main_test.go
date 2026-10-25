package main

import (
	"math"
	"testing"
)

// Chạy: go test ./02-flowcontrol/08-exercise-loops-and-functions
func TestSqrt(t *testing.T) {
	// Table-driven test — kiểu viết test idiomatic nhất của Go.
	for _, x := range []float64{0, 1, 2, 4, 9, 16, 100, 1e6, 0.25} {
		got := Sqrt(x)
		want := math.Sqrt(x)
		if math.Abs(got-want) > 1e-9 {
			t.Errorf("Sqrt(%v) = %v; want %v (sai lệch %v)", x, got, want, math.Abs(got-want))
		}
	}
}
