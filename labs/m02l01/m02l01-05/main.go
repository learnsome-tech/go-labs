# Concurrent Systems Programming with Go — lesson m02l01 — Pointers and Memory
# https://learnsome.tech/courses/go-course/watch?lesson=m02l01
# © LearnSome.tech
package main

import "fmt"

func describe(value *string) {
	if value == nil {
		fmt.Println("unset")
		return
	}
	fmt.Println("value:", *value)
}

func main() {
	var note *string
	describe(note)
	text := "ready"
	describe(&text)
}
