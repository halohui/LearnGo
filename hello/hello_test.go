package main

import "testing"

func add1(x, y int) int {
	return x + y
}

func BenckmarkAdd1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = add1(1, 2)
	}
}

func TestAdd1(t *testing.T) {
	if add1(1, 2) != 3 {
		t.FailNow()
	}
}
