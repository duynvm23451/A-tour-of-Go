package main

import (
	"fmt"
	"testing"
)

func TestIPAddrString(t *testing.T) {
	tests := []struct {
		ip   IPAddr
		want string
	}{
		{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
		{IPAddr{8, 8, 8, 8}, "8.8.8.8"},
		{IPAddr{0, 0, 0, 0}, "0.0.0.0"},
		{IPAddr{255, 254, 253, 252}, "255.254.253.252"},
	}

	for _, tt := range tests {
		// fmt.Sprint sẽ dùng method String nếu IPAddr cài đặt fmt.Stringer.
		if got := fmt.Sprint(tt.ip); got != tt.want {
			t.Errorf("fmt.Sprint(%v) = %q; want %q", [4]byte(tt.ip), got, tt.want)
		}
	}
}
