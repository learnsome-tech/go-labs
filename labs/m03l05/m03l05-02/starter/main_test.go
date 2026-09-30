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
