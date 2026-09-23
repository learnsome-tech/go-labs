# Concurrent Systems Programming with Go — lesson m03l04 — The Standard testing Package
# https://learnsome.tech/courses/go-course/watch?lesson=m03l04
# © LearnSome.tech
package main

import "testing"

func healthy(code int) bool { return code >= 200 && code < 300 }

func TestHealthy(t *testing.T) {
	if !healthy(204) {
		t.Error("success code should be healthy")
	}
	if healthy(500) {
		t.Error("server error should not be healthy")
	}
}
