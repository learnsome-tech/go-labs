# Concurrent Systems Programming with Go — lesson m02l03 — Structs and Methods
# https://learnsome.tech/courses/go-course/watch?lesson=m02l03
# © LearnSome.tech
package main

import "fmt"

type Target struct {
	Name string
	Port int
}

func (t Target) Address() string {
	return fmt.Sprintf("%s:%d", t.Name, t.Port)
}

func main() {
	t := Target{Name: "api", Port: 8080}
	fmt.Println(t.Address())
}
