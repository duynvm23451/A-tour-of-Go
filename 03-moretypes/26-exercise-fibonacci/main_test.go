package main

import "testing"

func TestFibonacci(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	f := fibonacci()
	for i, w := range want {
		if got := f(); got != w {
			t.Fatalf("lần gọi thứ %d trả về %d; want %d", i+1, got, w)
		}
	}

	// Mỗi lần gọi fibonacci() phải cho một closure ĐỘC LẬP, đếm lại từ đầu.
	g := fibonacci()
	if got := g(); got != 0 {
		t.Errorf("closure mới trả về %d ở lần gọi đầu; want 0", got)
	}
}
