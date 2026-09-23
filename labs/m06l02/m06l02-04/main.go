# Concurrent Systems Programming with Go — lesson m06l02 — Serving HTTP with net/http Server
# https://learnsome.tech/courses/go-course/watch?lesson=m06l02
# © LearnSome.tech
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	health(recorder, request)
	fmt.Println(recorder.Code, recorder.Body.String())
}
