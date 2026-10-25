// Package tree is a local stand-in for golang.org/x/tour/tree.
// Package tree là bản thay thế cục bộ cho golang.org/x/tour/tree,
// dùng cho bài tập "Equivalent Binary Trees".
package tree

import (
	"fmt"
	"math/rand"
	"strings"
)

// A Tree is a binary tree with integer values.
// Tree là cây nhị phân chứa số nguyên.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a random binary tree holding the values k, 2k, ..., 10k.
// New trả về một cây nhị phân ngẫu nhiên chứa các giá trị k, 2k, ..., 10k.
// Hình dạng cây khác nhau mỗi lần gọi, nhưng tập giá trị thì luôn giống nhau —
// đó chính là điểm mấu chốt của bài tập.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{Value: v}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

// String implements fmt.Stringer, so you can print a tree directly.
// String cài đặt fmt.Stringer, nhờ vậy in cây trực tiếp bằng fmt.Println được.
func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	var b strings.Builder
	b.WriteString("(")
	if t.Left != nil {
		fmt.Fprintf(&b, "%s ", t.Left)
	}
	fmt.Fprintf(&b, "%d", t.Value)
	if t.Right != nil {
		fmt.Fprintf(&b, " %s", t.Right)
	}
	b.WriteString(")")
	return b.String()
}
