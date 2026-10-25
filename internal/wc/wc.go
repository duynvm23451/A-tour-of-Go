// Package wc is a local stand-in for golang.org/x/tour/wc.
// Package wc là bản thay thế cục bộ cho golang.org/x/tour/wc,
// dùng để chấm bài tập WordCount.
package wc

import (
	"fmt"
	"sort"
	"strings"
)

// Test runs f against a handful of inputs and reports PASS/FAIL for each.
// Test chạy f với vài input mẫu rồi in PASS/FAIL cho từng trường hợp.
func Test(f func(string) map[string]int) {
	inputs := []string{
		"I am learning Go!",
		"a a a a a a",
		"",
		"the quick brown fox jumps over the lazy dog the end",
		"  spaces   everywhere  ",
	}

	for _, in := range inputs {
		got := f(in)
		want := reference(in)
		if equal(got, want) {
			fmt.Printf("PASS  f(%q) = %v\n", in, format(got))
		} else {
			fmt.Printf("FAIL  f(%q) = %v, want %v\n", in, format(got), format(want))
		}
	}
}

// reference is the obviously-correct implementation we compare against.
// reference là lời giải chuẩn dùng để đối chiếu.
func reference(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(s) {
		m[w]++
	}
	return m
}

func equal(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

// format prints a map with sorted keys so output is deterministic.
// format in map theo thứ tự key đã sắp xếp để output ổn định.
func format(m map[string]int) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	b.WriteString("map[")
	for i, k := range keys {
		if i > 0 {
			b.WriteByte(' ')
		}
		fmt.Fprintf(&b, "%s:%d", k, m[k])
	}
	b.WriteString("]")
	return b.String()
}
