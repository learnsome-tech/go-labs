# Concurrent Systems Programming with Go — lesson m03l05 — Writing Table Driven Tests
# https://learnsome.tech/courses/go-course/watch?lesson=m03l05
# © LearnSome.tech
package main

import "testing"

func TestHealthyTable(t *testing.T) {
	cases := []struct { name string; code int; want bool }{
		{"ok", 200, true},
		{"redirect", 302, false},
		{"error", 500, false},
	}
	for _, tc := range cases {
		if got := healthy(tc.code); got != tc.want {
			t.Errorf("%s: got %v", tc.name, got)
		}
		}
}
