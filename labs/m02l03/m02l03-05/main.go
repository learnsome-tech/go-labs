# Concurrent Systems Programming with Go — lesson m02l03 — Structs and Methods
# https://learnsome.tech/courses/go-course/watch?lesson=m02l03
# © LearnSome.tech
package main

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) Add() {
	c.Value++
}

func main() {
	c := Counter{}
	c.Add()
	c.Add()
	fmt.Println(c.Value)
}
