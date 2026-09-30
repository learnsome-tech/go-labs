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
