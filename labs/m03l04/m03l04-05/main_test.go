# Concurrent Systems Programming with Go — lesson m03l04 — The Standard testing Package
# https://learnsome.tech/courses/go-course/watch?lesson=m03l04
# © LearnSome.tech
package main

import "testing"

func parsePort(text string) (int, error) {
	if text == "" { return 0, errEmpty }
	return 8080, nil
}

var errEmpty = testing.Errs()

func TestParsePortError(t *testing.T) {
	if _, err := parsePort(""); err == nil {
		t.Fatal("empty text should fail")
	}
}
