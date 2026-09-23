# Concurrent Systems Programming with Go — lesson m02l03 — Structs and Methods
# https://learnsome.tech/courses/go-course/watch?lesson=m02l03
# © LearnSome.tech
package main

import "fmt"

type Target struct {
	Name string
	Port int
}

func main() {
	target := Target{Name: "api", Port: 8080}
	fmt.Println(target.Name, target.Port)
}
