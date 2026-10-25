package main

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		in   string
		want map[string]int
	}{
		{"", map[string]int{}},
		{"a", map[string]int{"a": 1}},
		{"a a a", map[string]int{"a": 3}},
		{"I am learning Go!", map[string]int{"I": 1, "am": 1, "learning": 1, "Go!": 1}},
		{"  nhieu   khoang   trang  ", map[string]int{"nhieu": 1, "khoang": 1, "trang": 1}},
	}

	for _, tt := range tests {
		if got := WordCount(tt.in); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("WordCount(%q) = %v; want %v", tt.in, got, tt.want)
		}
	}
}
