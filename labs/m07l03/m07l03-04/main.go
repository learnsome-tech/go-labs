# Concurrent Systems Programming with Go — lesson m07l03 — Handling Timeouts and Errors
# https://learnsome.tech/courses/go-course/watch?lesson=m07l03
# © LearnSome.tech
package main

import "fmt"

type Result struct { Name string; Err error }

func main() {
	result := Result{Name: "api", Err: fmt.Errorf("timeout")}
	fmt.Println(result.Name, result.Err)
}
