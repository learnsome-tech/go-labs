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
