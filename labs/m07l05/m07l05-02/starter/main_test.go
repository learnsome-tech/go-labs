package main

import "testing"

func TestHealthyResult(t *testing.T) {
	result := Result{Name: "api", Healthy: true}
	if !result.Healthy {
		t.Fatal("api should be healthy")
	}
}
