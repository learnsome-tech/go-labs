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
