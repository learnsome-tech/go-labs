# Concurrent Systems Programming with Go — lesson m04l02 — Channels and Synchronization
# https://learnsome.tech/courses/go-course/watch?lesson=m04l02
# © LearnSome.tech
package main

import "fmt"

func main() {
	values := make(chan int, 2)
	values <- 3
	values <- 5
	close(values)
	for value := range values {
		fmt.Println(value)
	}
	fmt.Println("done")
}
