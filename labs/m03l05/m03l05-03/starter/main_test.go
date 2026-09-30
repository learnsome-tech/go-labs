package main

import "testing"

func TestHealthySubtests(t *testing.T) {
	cases := map[string]int{"ok": 200, "error": 500}
	for name, code := range cases {
		t.Run(name, func(t *testing.T) {
			if !healthy(code) && name == "ok" { t.Fatal("expected healthy") }
		})
	}
}
