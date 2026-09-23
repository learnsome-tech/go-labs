# Concurrent Systems Programming with Go — lesson m05l05 — Calling APIs with net/http Client
# https://learnsome.tech/courses/go-course/watch?lesson=m05l05
# © LearnSome.tech
package main

import (
	"fmt"
	"net/http"
)

func main() {
	request, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:1", nil)
	_, err := http.DefaultClient.Do(request)
	fmt.Println(err != nil)
}
