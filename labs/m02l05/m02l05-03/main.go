# Concurrent Systems Programming with Go — lesson m02l05 — Type Assertions and the Empty Interface
# https://learnsome.tech/courses/go-course/watch?lesson=m02l05
# © LearnSome.tech
package main

import "fmt"

func describe(value any) {
	switch item := value.(type) {
	case string:
		fmt.Println("text", item)
	case int:
		fmt.Println("number", item)
	default:
		fmt.Println("other")
	}
}

func main() {
	describe("api")
	describe(3)
	describe(true)
}
