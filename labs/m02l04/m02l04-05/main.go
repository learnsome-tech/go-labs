# Concurrent Systems Programming with Go — lesson m02l04 — Interfaces and Why They Are Implicit
# https://learnsome.tech/courses/go-course/watch?lesson=m02l04
# © LearnSome.tech
package main

import "fmt"

type Reader interface {
	Read() string
}

type Fixture struct{}

func (Fixture) Read() string { return "fixture data" }

func load(r Reader) string {
	return r.Read()
}

func main() {
	fmt.Println(load(Fixture{}))
}
