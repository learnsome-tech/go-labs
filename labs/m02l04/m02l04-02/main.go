# Concurrent Systems Programming with Go — lesson m02l04 — Interfaces and Why They Are Implicit
# https://learnsome.tech/courses/go-course/watch?lesson=m02l04
# © LearnSome.tech
package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Robot struct{}

func (Robot) Speak() string { return "ready" }

func announce(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	announce(Robot{})
}
