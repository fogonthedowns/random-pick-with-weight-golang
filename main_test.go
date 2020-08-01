package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	w := []int{1, 0}
	obj := Constructor(w)
	out := obj.PickIndex()

	if out != 0 {
		t.Errorf("got %#v, want %#v", out, 0)
	}
}
