# Concurrent Systems Programming with Go — lesson m05l05 — Calling APIs with net/http Client
# https://learnsome.tech/courses/go-course/watch?lesson=m05l05
# © LearnSome.tech
package main

import (
	"fmt"
	"io"
	"net/http"
)

type fake struct{}
func (fake) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 200, Body: io.NopCloser(nil)}, nil
}

func main() {
	client := &http.Client{Transport: fake{}}
	response, err := client.Get("http://example.invalid")
	fmt.Println(response.StatusCode, err)
	response.Body.Close()
}
