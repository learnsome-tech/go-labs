# Concurrent Systems Programming with Go — lesson m04l02 — Channels and Synchronization
# https://learnsome.tech/courses/go-course/watch?lesson=m04l02
# © LearnSome.tech
package main

import "fmt"

func main() {
	values := make(chan string)
	go func() { values <- "ready" }()
	value := <-values
	fmt.Println(value)
}
