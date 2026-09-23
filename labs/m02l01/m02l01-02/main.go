# Concurrent Systems Programming with Go — lesson m02l01 — Pointers and Memory
# https://learnsome.tech/courses/go-course/watch?lesson=m02l01
# © LearnSome.tech
package main

import "fmt"

func main() {
	name := "api"
	p := &name
	fmt.Println(*p)
	*p = "worker"
	fmt.Println(name)
}
