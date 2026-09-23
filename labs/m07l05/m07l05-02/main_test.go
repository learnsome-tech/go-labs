# Concurrent Systems Programming with Go — lesson m07l05 — Containerizing and Testing the Final Build
# https://learnsome.tech/courses/go-course/watch?lesson=m07l05
# © LearnSome.tech
package main

import "testing"

func TestHealthyResult(t *testing.T) {
	result := Result{Name: "api", Healthy: true}
	if !result.Healthy {
		t.Fatal("api should be healthy")
	}
}
