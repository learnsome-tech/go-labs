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
