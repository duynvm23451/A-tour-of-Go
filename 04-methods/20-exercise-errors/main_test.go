package main

import (
	"fmt"
	"math"
	"testing"
)

func TestSqrtPositive(t *testing.T) {
	for _, x := range []float64{0, 1, 2, 4, 9, 16, 100} {
		got, err := Sqrt(x)
		if err != nil {
			t.Errorf("Sqrt(%v) trả về lỗi %v; want nil", x, err)
			continue
		}
		if math.Abs(got-math.Sqrt(x)) > 1e-9 {
			t.Errorf("Sqrt(%v) = %v; want %v", x, got, math.Sqrt(x))
		}
	}
}

func TestSqrtNegative(t *testing.T) {
	_, err := Sqrt(-2)
	if err == nil {
		t.Fatal("Sqrt(-2) trả về nil error; want ErrNegativeSqrt")
	}
	if want := "cannot Sqrt negative number: -2"; err.Error() != want {
		t.Errorf("err.Error() = %q; want %q", err.Error(), want)
	}
	// So sánh tên kiểu bằng chuỗi thay vì type assertion, để file test này
	// vẫn biên dịch được khi ErrNegativeSqrt chưa có method Error.
	if got := fmt.Sprintf("%T", err); got != "main.ErrNegativeSqrt" {
		t.Errorf("err có kiểu %s; want main.ErrNegativeSqrt", got)
	}
}
