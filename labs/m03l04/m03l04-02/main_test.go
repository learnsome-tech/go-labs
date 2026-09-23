# Concurrent Systems Programming with Go — lesson m03l04 — The Standard testing Package
# https://learnsome.tech/courses/go-course/watch?lesson=m03l04
# © LearnSome.tech
package main

import "testing"

func double(value int) int { return value * 2 }

func TestDouble(t *testing.T) {
	if got := double(4); got != 8 {
		t.Fatalf("double returned %d", got)
	}
}
