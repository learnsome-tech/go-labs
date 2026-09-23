# Concurrent Systems Programming with Go — lesson m02l01 — Pointers and Memory
# https://learnsome.tech/courses/go-course/watch?lesson=m02l01
# © LearnSome.tech
package main

import "fmt"

func addTag(label *string) {
	*label = *label + "-ready"
}

func main() {
	status := "service"
	addTag(&status)
	fmt.Println(status)
}
