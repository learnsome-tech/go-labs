# Concurrent Systems Programming with Go — lesson m02l04 — Interfaces and Why They Are Implicit
# https://learnsome.tech/courses/go-course/watch?lesson=m02l04
# © LearnSome.tech
package main

import "fmt"

type Reporter interface {
	Report() string
}

type Healthy struct{}
type Failed struct{}

func (Healthy) Report() string { return "healthy" }
func (Failed) Report() string { return "failed" }

func printReport(r Reporter) {
	fmt.Println(r.Report())
}

func main() {
	printReport(Healthy{})
	printReport(Failed{})
}
