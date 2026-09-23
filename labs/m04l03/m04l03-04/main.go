# Concurrent Systems Programming with Go — lesson m04l03 — The select Statement
# https://learnsome.tech/courses/go-course/watch?lesson=m04l03
# © LearnSome.tech
package main

import "fmt"

func main() {
	updates := make(chan string, 1)
	select {
	case value := <-updates:
		fmt.Println(value)
	default:
		fmt.Println("no update")
	}
}
