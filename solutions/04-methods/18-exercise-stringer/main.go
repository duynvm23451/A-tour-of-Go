// Lời giải: Stringers (IPAddr)
package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

// String cài đặt fmt.Stringer, nên fmt.Printf("%v", ip) sẽ dùng nó.
//
// Dùng value receiver (không phải *IPAddr): nhờ vậy cả IPAddr lẫn *IPAddr
// đều thỏa mãn fmt.Stringer, và map[string]IPAddr in ra đúng như mong đợi.
func (ip IPAddr) String() string {
	parts := make([]string, len(ip))
	for i, b := range ip {
		parts[i] = fmt.Sprint(b) // b là byte -> in ra dạng số
	}
	return strings.Join(parts, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
