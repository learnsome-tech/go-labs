# Concurrent Systems Programming with Go — lesson m06l02 — Serving HTTP with net/http Server
# https://learnsome.tech/courses/go-course/watch?lesson=m06l02
# © LearnSome.tech
package main

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "ok")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)
	fmt.Println(mux != nil)
}
