# Concurrent Systems Programming with Go — lesson m02l05 — Type Assertions and the Empty Interface
# https://learnsome.tech/courses/go-course/watch?lesson=m02l05
# © LearnSome.tech
package main

import "fmt"

func show(value any) {
	text, ok := value.(string)
	if !ok {
		fmt.Println("not text")
		return
	}
	fmt.Println("text:", text)
}

func main() {
	show("ready")
	show(42)
}
