package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//!-

func (t *tree) String() string {
	values := make([]int, 0)
	values = appendValues(values, t)
	if len(values) == 0 {
		return "tree[]"
	}

	buffer := &bytes.Buffer{}
	fmt.Fprintf(buffer, "tree[%d", values[0])
	for _, v := range values[1:] {
		fmt.Fprintf(buffer, ", %d", v)
	}
	fmt.Fprintf(buffer, "]")
	return buffer.String()
}
