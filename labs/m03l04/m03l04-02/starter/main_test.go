package main

import "testing"

func double(value int) int { return value * 2 }

func TestDouble(t *testing.T) {
	if got := double(4); got != 8 {
		t.Fatalf("double returned %d", got)
	}
}
