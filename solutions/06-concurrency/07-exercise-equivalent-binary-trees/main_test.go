package main

import (
	"testing"

	"tour/internal/tree"
)

func TestWalk(t *testing.T) {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var got []int
	for v := range ch { // dựa vào việc Walk close channel khi xong
		got = append(got, v)
	}

	if len(got) != len(want) {
		t.Fatalf("Walk gửi %d giá trị (%v); want %d", len(got), got, len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("Walk = %v; want %v", got, want)
		}
	}
}

func TestSame(t *testing.T) {
	if !Same(tree.New(1), tree.New(1)) {
		t.Error("Same(tree.New(1), tree.New(1)) = false; want true")
	}
	if Same(tree.New(1), tree.New(2)) {
		t.Error("Same(tree.New(1), tree.New(2)) = true; want false")
	}
}
