package main

import "testing"

func TestPortTable(t *testing.T) {
	cases := []struct { name, input string; want int; bad bool }{
		{"empty", "", 0, true},
		{"service", "api", 8080, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parsePort(tc.input)
			bad := (err != nil) != tc.bad
			if bad || (!tc.bad && got != tc.want) {
				t.Fatal("contract mismatch")
			}
		})
	}
}
