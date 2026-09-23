# Concurrent Systems Programming with Go — lesson m07l01 — Project Setup and Configuration
# https://learnsome.tech/courses/go-course/watch?lesson=m07l01
# © LearnSome.tech
package main

import "fmt"

type Target struct { Name string; URL string }

func main() {
	targets := []Target{{"api", "https://api"}, {"web", "https://web"}}
	for _, target := range targets {
		fmt.Println(target.Name, target.URL)
	}
}
