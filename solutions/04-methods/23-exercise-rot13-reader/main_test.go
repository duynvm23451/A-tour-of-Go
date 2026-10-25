package main

import (
	"io"
	"strings"
	"testing"
)

func TestRot13Reader(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"Lbh penpxrq gur pbqr!", "You cracked the code!"},
		{"abc XYZ", "nop KLM"},
		{"12345 !@#", "12345 !@#"}, // ký tự không phải chữ cái giữ nguyên
		{"", ""},
	}

	for _, tt := range tests {
		r := &rot13Reader{strings.NewReader(tt.in)}
		got, err := io.ReadAll(r)
		if err != nil {
			t.Fatalf("ReadAll(%q) lỗi: %v", tt.in, err)
		}
		if string(got) != tt.want {
			t.Errorf("rot13(%q) = %q; want %q", tt.in, got, tt.want)
		}
	}
}
