# Concurrent Systems Programming with Go — lesson m03l03 — Panics and Recover
# https://learnsome.tech/courses/go-course/watch?lesson=m03l03
# © LearnSome.tech
package main

import "fmt"

func work() {
	defer fmt.Println("cleanup")
	defer func() {
		if recover() != nil {
			fmt.Println("recovered")
		}
	}()
	panic("failure")
}

func main() {
	work()
	fmt.Println("done")
}
