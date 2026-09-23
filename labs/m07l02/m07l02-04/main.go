# Concurrent Systems Programming with Go — lesson m07l02 — Executing Concurrent Checks
# https://learnsome.tech/courses/go-course/watch?lesson=m07l02
# © LearnSome.tech
package main

import "fmt"

type Result struct { Name string; Healthy bool }

func main() {
	results := []Result{{"api", true}, {"web", false}}
	for _, result := range results {
		fmt.Println(result.Name, result.Healthy)
	}
}
